package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
	"github.com/shutej/gopherjs-test/service"
)

func main() {
	conn, _ := websocket.Dial("ws://127.0.0.1:8000/")
	client := rpcplus.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	quotient := service.Quotient{}
	_ = client.Call("Arith.Divide", service.Args{A: 2, B: 4}, &quotient)
	js.Global.Get("document").Call("write", "Hello World!")
}
