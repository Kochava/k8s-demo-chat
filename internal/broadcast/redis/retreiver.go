package redisbroadcast

import (
	"io"

	redis "github.com/go-redis/redis"
)

// Retriever receives messages from Redis
type Retriever struct {
	Writer io.Writer
	PubSub *redis.PubSub
}

// Retrieve reads messages from a redis channel and supplies them to Writer
func (retrieve *Retriever) Retrieve() {
	for message := range retrieve.PubSub.Channel() {
		retrieve.Writer.Write([]byte(message.Payload))
	}
}
