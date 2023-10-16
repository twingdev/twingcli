package server

import (
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p/core/network"
	pb "github.com/twingdev/twingcli/server/pb"
	"io"
	"log"
	"sync"
)

const healthRequest = "/health/req/1.0"
const healthResponse = "/health/res/1.0"

type HealthProtocol struct {
	node     *Node
	mu       sync.Mutex
	requests map[string]*pb.HealthRequest
	done     chan bool
}

func NewHealthProtocol(node *Node, done chan bool) *HealthProtocol {
	h := &HealthProtocol{node: node, requests: make(map[string]*pb.HealthRequest), done: done}
	node.Host.SetStreamHandler(healthRequest, h.onHealthReq)
	return h
}

func (h *HealthProtocol) onHealthReq(s network.Stream) {
	data := &pb.HealthRequest{}
	buf, err := io.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}
	s.Close()

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}
}
