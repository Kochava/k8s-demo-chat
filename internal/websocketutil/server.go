package websocketutil

import (
	"context"
	"net/http"

	"golang.org/x/net/websocket"
)

// Server defines an HTTP style server for WebSocket connections
type Server struct {
	Addr       string
	HandleFunc websocket.Handler

	httpServer *http.Server
}

// ListenAndServe accepts new connections and routes them to the handler
func (server *Server) ListenAndServe() error {
	var (
		err error
		mux = http.NewServeMux()
	)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	mux.Handle("/chat", server.HandleFunc)

	server.httpServer = &http.Server{
		Addr:    server.Addr,
		Handler: mux,
	}

	if err = server.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Shutdown closes the listener to stop new connections
func (server *Server) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
