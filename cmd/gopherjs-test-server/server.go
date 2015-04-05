package main

import (
	"net/http"

	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
	"golang.org/x/net/websocket"
)

func defaultFactory() interface{} {
	return nil
}

type Server struct {
	server  *rpcplus.Server
	factory func() interface{}
}

type Option func(server *Server)

func ContextFactory(factory func() interface{}) Option {
	return func(self *Server) {
		self.factory = factory
	}
}

func NewServer(server *rpcplus.Server, options ...Option) *Server {
	if server == nil {
		server = rpcplus.NewServer()
	}

	self := &Server{
		factory: defaultFactory,
		server:  server,
	}

	for _, option := range options {
		option(self)
	}

	return self
}

func (self *Server) Handler() http.Handler {
	return websocket.Handler(func(conn *websocket.Conn) {
		codec := jsonrpc.NewServerCodec(conn)
		context := self.factory()
		self.server.ServeCodecWithContext(codec, context)
	})
}
