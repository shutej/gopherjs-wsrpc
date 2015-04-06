package main

import (
	"fmt"
	"log"

	"github.com/gopherjs/gopherjs/js"
	"github.com/shutej/gopherjs-wsrpc/service"
	"github.com/shutej/wsrpc/client"
)

//go:generate gopherjs build -m main.go
func main() {
	c, err := client.New("ws://127.0.0.1:8000/jsonrpc")
	if err != nil {
		log.Fatal(err)
	}

	quotient := service.Quotient{}
	err = c.Call("Arith.Divide", service.Args{A: 9, B: 4}, &quotient)
	js.Global.Get("document").Call("write", fmt.Sprintf("%#v", quotient))
}
