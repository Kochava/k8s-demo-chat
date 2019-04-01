package broadcast

import (
	"net"

	"github.com/Kochava/k8s-demo-chat/internal/tcputil"
)

var (
	_ tcputil.Handler = &TCPHandlerProxy{}
)

// TCPHandlerProxy proxies a TCPHandler to a ReadWriteHandler
type TCPHandlerProxy struct {
	ReadWriteHandler *ReadWriteHandler
}

// Handle processes a tcp connection
func (handler *TCPHandlerProxy) Handle(conn net.Conn) {
	handler.ReadWriteHandler.Handle(conn)
}
