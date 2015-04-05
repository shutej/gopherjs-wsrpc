package wsrpc

import (
	"net/http"

	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
	"golang.org/x/net/websocket"
)

func defaultFactory() interface{} {
	return nil
}

type server struct {
	server  *rpcplus.Server
	factory func() interface{}
}

type Option func(server *server)

func ContextFactory(factory func() interface{}) Option {
	return func(self *server) {
		self.factory = factory
	}
}

func Handler(s *rpcplus.Server, options ...Option) http.Handler {
	if s == nil {
		s = rpcplus.NewServer()
	}

	self := &server{
		factory: defaultFactory,
		server:  s,
	}

	for _, option := range options {
		option(self)
	}

	return websocket.Handler(func(conn *websocket.Conn) {
		codec := jsonrpc.NewServerCodec(conn)
		context := self.factory()
		self.server.ServeCodecWithContext(codec, context)
	})
}
