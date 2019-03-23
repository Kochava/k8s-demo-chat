package redisbroadcast

import "io"

var (
	_ io.Writer = &Writer{}
)

// Writer stores messages to a message channel
type Writer struct {
	MessageChannel chan<- []byte
}

func (writer *Writer) Write(message []byte) (int, error) {
	writer.MessageChannel <- message
	return len(message), nil
}
