package redisbroadcast

import "io"

var (
	_ io.Writer = &Writer{}
)

// Writer provides go routine support for Publisher.
//
// The go-redis/redis package's pub/sub doesn't support
// concurrency. This struct does support concurrency and
// writes to a channel which Publisher can read from
type Writer struct {
	MessageChannel chan<- []byte
}

// Write stores messages to a message channel
func (writer *Writer) Write(message []byte) (int, error) {
	writer.MessageChannel <- message
	return len(message), nil
}
