package tcputil

import "io"

var (
	_ io.Writer = &TCPWriter{}
)

func NewTCPWriter(connWriter io.Writer) io.Writer {
	return &TCPWriter{connWriter: connWriter}
}

type TCPWriter struct {
	connWriter io.Writer
}

func (tcpWriter *TCPWriter) Write(msg []byte) (int, error) {
	var (
		connInt   int
		connError error
	)

	connInt, connError = tcpWriter.connWriter.Write(msg)
	tcpWriter.connWriter.Write([]byte("\n"))

	return connInt, connError
}
