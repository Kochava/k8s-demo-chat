package websocketutil

import "golang.org/x/net/websocket"

// Handler defines a common connection handler
type Handler interface {
	Handle(*websocket.Conn)
}
