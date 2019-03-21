package websocketutil

import (
	"context"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	Addr       string
	HandleFunc websocket.Handler

	httpServer *http.Server
}

func (server *Server) ListenAndServe() error {
	var (
		err error
		mux = http.NewServeMux()
	)

	mux.Handle("/", server.HandleFunc)

	server.httpServer = &http.Server{
		Addr:    server.Addr,
		Handler: mux,
	}

	if err = server.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
