package build

import redis "github.com/go-redis/redis"

// Redis uses config to build a redis client
func Redis(config *Config) (*redis.Client, error) {
	var (
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.RedisAddr,
		})
	)

	return redisClient, nil
}
