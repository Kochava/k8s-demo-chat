package broadcast

import (
	"log"

	"git.dev.kochava.com/notter/distchat/websocketutil"
	"golang.org/x/net/websocket"
)

var (
	_ websocketutil.Handler = &WebsocketHandlerProxy{}
)

type WebsocketHandlerProxy struct {
	ReadWriteHandler *ReadWriteHandler
}

func (handler *WebsocketHandlerProxy) Handle(conn *websocket.Conn) {
	log.Println("Handling the connection")
	handler.ReadWriteHandler.Handle(conn)
}
