package redisbroadcast

import (
	"log"

	redis "github.com/go-redis/redis"
)

// Publisher publishes messages to Redis
type Publisher struct {
	MessageChannel <-chan []byte
	RedisChannel   string
	RedisClient    *redis.Client
}

// Publish reads incomming messages and publishes them to a redis channel
func (publisher *Publisher) Publish() {
	for message := range publisher.MessageChannel {
		var intCmd = publisher.RedisClient.Publish(
			publisher.RedisChannel,
			string(message),
		)

		if intCmd.Err() != nil {
			log.Println("unable to store message:", intCmd.Err().Error())
		}
	}
}
