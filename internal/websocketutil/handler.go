package websocketutil

import "golang.org/x/net/websocket"

type Handler interface {
	Handle(*websocket.Conn)
}
