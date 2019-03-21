package redisbroadcast

import (
	"io"

	redis "github.com/go-redis/redis"
)

type Retriever struct {
	Writer io.Writer
	PubSub *redis.PubSub
}

func (retrieve *Retriever) Retrieve() {
	for message := range retrieve.PubSub.Channel() {
		retrieve.Writer.Write([]byte(message.Payload))
	}
}
