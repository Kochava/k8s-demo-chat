package tcputil

import "io"

var (
	_ io.Writer = &TCPWriter{}
)

// NewTCPWriter constructs a new TCPWriter
func NewTCPWriter(connWriter io.Writer) io.Writer {
	return &TCPWriter{connWriter: connWriter}
}

// TCPWriter wraps an io.Writer for communication of TCP
type TCPWriter struct {
	connWriter io.Writer
}

// Write adds a newline after the message to send the message to a TCP connection
func (tcpWriter *TCPWriter) Write(msg []byte) (int, error) {
	var (
		connInt   int
		connError error
	)

	connInt, connError = tcpWriter.connWriter.Write(msg)
	tcpWriter.connWriter.Write([]byte("\n"))

	return connInt, connError
}
