package broadcast

import (
	"github.com/Kochava/k8s-demo-chat/internal/websocketutil"
	"golang.org/x/net/websocket"
)

var (
	_ websocketutil.Handler = &WebsocketHandlerProxy{}
)

// WebsocketHandlerProxy proxies a WebSocketHandler to a ReadWriteHandler
type WebsocketHandlerProxy struct {
	ReadWriteHandler *ReadWriteHandler
}

// Handle processes a websocket connection
func (handler *WebsocketHandlerProxy) Handle(conn *websocket.Conn) {
	handler.ReadWriteHandler.Handle(conn)
}
