package tcputil

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	Addr    string
	Handler Handler

	waitGroup *sync.WaitGroup
	listener  net.Listener
}

func (server *Server) ListenAndServe() error {
	var (
		err error
	)

	server.waitGroup = &sync.WaitGroup{}

	if server.listener, err = net.Listen("tcp", server.Addr); err != nil {
		return err
	}

	for {
		var (
			err        error
			connection net.Conn
		)

		if connection, err = server.listener.Accept(); err != nil {
			log.Println("Unable to accept connection:", err.Error())
			break
		}

		server.waitGroup.Add(1)
		go server.handleConnection(connection)
	}

	server.waitGroup.Wait()
	return nil
}

func (server *Server) handleConnection(connection net.Conn) {
	defer server.waitGroup.Done()
	server.Handler.Handle(connection)
}

func (server *Server) Shutdown() error {
	return server.listener.Close()
}
