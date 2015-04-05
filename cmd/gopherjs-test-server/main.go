package main

import (
	"errors"
	"net/http"

	httpgzip "github.com/daaku/go.httpgzip"
	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/gopherjs-test/service"
	"github.com/shutej/gopherjs-test/wsrpc"
)

type Arith int

func (t *Arith) Multiply(args *service.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *service.Args, quo *service.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	rpc := rpcplus.NewServer()
	rpc.Register(new(Arith))

	http.Handle("/jsonrpc", wsrpc.Handler(rpc))
	http.Handle("/", http.StripPrefix("/", httpgzip.NewHandler(http.FileServer(http.Dir("static")))))

	http.ListenAndServe(":8000", nil)
}
