package redis

import (
	redislib "github.com/go-redis/redis"
)

// BuildGoRedis uses the supplied configuration to build a redis connection
func BuildGoRedis(config *Config) (*redislib.Client, error) {
	var (
		redisClient = redislib.NewClient(&redislib.Options{
			Addr: config.Addr,
		})
	)

	return redisClient, nil
}
