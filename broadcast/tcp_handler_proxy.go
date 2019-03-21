package broadcast

import (
	"net"

	"git.dev.kochava.com/notter/distchat/tcputil"
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
