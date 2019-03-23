package tcputil

import (
	"net"
)

type Handler interface {
	Handle(net.Conn)
}
