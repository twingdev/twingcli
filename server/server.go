package server

import (
	"context"
	"crypto"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	crypt "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/twingdev/twingcli/common/fs"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Node struct {
	ID         peer.ID
	Host       host.Host
	GrpcServer grpc.Server
	Reader     chan *fs.StreamIO
}

type IServer interface {
	Build(out host.Host)
	//Bootstrap() error
}

func NewNode(ctx context.Context) {
	n := &Node{}
	n.Build(ctx, n.Host)

}

var _ crypto.PrivateKey

func (n *Node) Build(ctx context.Context, out host.Host) {

	var idht *dht.IpfsDHT

	cmgr, err := connmgr.NewConnManager(
		100, // Lowwater
		400, // HighWater,
		connmgr.WithGracePeriod(time.Minute),
	)
	if err != nil {
		panic(err)
	}

	// Set your own keypair
	priv, _, err := crypt.GenerateKeyPair(
		crypt.Secp256k1, // Select your key type. Ed25519 are nice short
		-1,              // Select key length when possible (i.e. RSA).
	)
	if err != nil {
		panic(err)
	}

	out, err = libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(priv),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9000",      // regular tcp connections
			"/ip4/0.0.0.0/udp/9000/quic", // a UDP endpoint for the QUIC transport
		),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(cmgr),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err = dht.New(ctx, h)
			return idht, err
		}),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)
	if err != nil {
		panic(err)
	}

	n.ID = out.ID()

	defer out.Close()
	for _, addr := range dht.DefaultBootstrapPeers {
		pi, _ := peer.AddrInfoFromP2pAddr(addr)
		if pi.ID != n.ID {

			// We ignore errors as some bootstrap peers may be down
			// and that is fine.
			out.Connect(ctx, *pi)
			log.Printf("peer connected %v", pi.ID)

		}
	}

}
