package build

import (
	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	redisbroadcast "github.com/Kochava/k8s-demo-chat/internal/broadcast/redis"
	buildredis "github.com/Kochava/k8s-demo-chat/internal/build/redis"
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

	if redisClient, err = buildredis.BuildGoRedis(config.Redis); err != nil {
		return nil, err
	}

	pubsub = redisClient.Subscribe(config.Redis.Channel)

	retriever = redisbroadcast.Retriever{
		PubSub: pubsub,
		Writer: broadcastWriter,
	}
	go retriever.Retrieve()

	publisher = redisbroadcast.Publisher{
		RedisChannel:   config.Redis.Channel,
		MessageChannel: messagesToStore,
		RedisClient:    redisClient,
	}
	go publisher.Publish()

	return readWriteHandler, nil
}
