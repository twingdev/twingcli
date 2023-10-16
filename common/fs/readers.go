package fs

import (
	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type StreamIO struct {
	r *ggio.Reader
	w *ggio.Writer
}

func (io *StreamIO) sendProtoMessage(id peer.ID, p protocol.ID, data proto.Message) bool {

}
