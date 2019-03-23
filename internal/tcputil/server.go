package tcputil

import (
	"context"
	"log"
	"net"
)

type Server struct {
	Addr    string
	Handler Handler

	listener net.Listener
}

func (server *Server) ListenAndServe() error {
	var (
		err error
	)

	if server.listener, err = net.Listen("tcp", server.Addr); err != nil {
		return err
	}

	for {
		var (
			err        error
			connection net.Conn
		)

		if connection, err = server.listener.Accept(); err != nil {
			log.Println("Unable to accept connection")
			continue
		}

		go server.Handler.Handle(connection)
	}
}

func (server *Server) Shutdown(_ context.Context) error {
	return server.listener.Close()
}
