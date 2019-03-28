package broadcast

import (
	"net"

	"github.com/Kochava/k8s-demo-chat/internal/tcputil"
)

var (
	_ tcputil.Handler = &TCPHandlerProxy{}
)

type TCPHandlerProxy struct {
	ReadWriteHandler *ReadWriteHandler
}

func (handler *TCPHandlerProxy) Handle(conn net.Conn) {
	handler.ReadWriteHandler.Handle(conn)
}
