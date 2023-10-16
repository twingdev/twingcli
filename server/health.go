package server

import (
	"github.com/libp2p/go-libp2p/core/network"

	"sync"
)

const healthRequest = "/health/req/1.0"
const healthResponse = "/health/res/1.0"

type HealthProtocol struct {
	node     *Node
	mu       sync.Mutex
	requests map[string]interface{}
	done     chan bool
}

func NewHealthProtocol(node *Node, done chan bool) *HealthProtocol {
	h := &HealthProtocol{node: node, requests: make(map[string]interface{}), done: done}
	node.Host.SetStreamHandler(healthRequest, h.onHealthReq)
	return h
}

func (h *HealthProtocol) onHealthReq(s network.Stream) {

}
