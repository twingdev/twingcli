package protocol

import (
	"encoding/json"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"os"
	"reflect"
)

const PROTOCOL_ID = protocol.ID("/twing/1.0/")

const ROUTE_CONFIG = "./configs/router.json"

type HandlerRepo struct {
	RouteName   string `json:"route_name"`
	HandlerName string `json:"handler_name"`
}

var streamHandler map[string]func(s network.Stream)

func initRouter() {
	streamHandler = make(map[string]func(s network.Stream))
	var hdl []HandlerRepo

	ReadJSON(ROUTE_CONFIG, hdl)

}

func ReadJSON(filename string, out interface{}) {
	data, _ := os.ReadFile(filename)
	_ = json.Unmarshal(data, &out)
}

func ExecuteHandler(funcName string, vals ...reflect.Value) {

	f := reflect.ValueOf(streamHandler[funcName])
	f.Elem().Call(vals)

}
