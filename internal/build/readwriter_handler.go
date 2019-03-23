package build

import (
	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/broadcast/redis"
	redis "github.com/go-redis/redis"
)

func GetReadWriterHandler(config *Config) (*broadcast.ReadWriteHandler, error) {
	var (
		err error

		messagesToStore = make(chan []byte, 10)

		redisClient *redis.Client
		pubsub      *redis.PubSub

		retriever redisbroadcast.Retriever
		publisher redisbroadcast.Publisher

		byteChannelWriter = &redisbroadcast.Writer{
			MessageChannel: messagesToStore,
		}

		broadcastWriter = broadcast.NewWriter()

		readWriteHandler = &broadcast.ReadWriteHandler{
			Writer:        byteChannelWriter,
			WriterStorage: broadcastWriter,
		}
	)

	if readWriteHandler.InputValidator, err = broadcast.NewJSONSchemaValidator(config.JSONValidationSchemaPath); err != nil {
		return nil, err
	}

	if redisClient, err = Redis(config); err != nil {
		return nil, err
	}

	pubsub = redisClient.Subscribe(config.RedisChannel)

	retriever = redisbroadcast.Retriever{
		PubSub: pubsub,
		Writer: broadcastWriter,
	}

	go retriever.Retrieve()

	publisher = redisbroadcast.Publisher{
		MessageChannel: messagesToStore,
		RedisChannel:   config.RedisChannel,
		RedisClient:    redisClient,
	}

	go publisher.Publish()

	return readWriteHandler, nil
}
