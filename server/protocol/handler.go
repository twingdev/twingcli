package protocol

import (
	"github.com/libp2p/go-libp2p/core/protocol"
	"reflect"
)

const PROTOCOL_ID = protocol.ID("/twing/1.0/")

type HandlerRepo struct {
	RouteName   string `json:"route_name"`
	HandlerName string `json:"handler_name"`
}

func Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
