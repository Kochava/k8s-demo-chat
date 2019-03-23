package broadcast

import (
	"github.com/Kochava/k8s-demo-chat/internal/websocketutil"
	"golang.org/x/net/websocket"
)

var (
	_ websocketutil.Handler = &WebsocketHandlerProxy{}
)

type WebsocketHandlerProxy struct {
	ReadWriteHandler *ReadWriteHandler
}

func (handler *WebsocketHandlerProxy) Handle(conn *websocket.Conn) {
	handler.ReadWriteHandler.Handle(conn)
}
