package redis

import (
	redislib "github.com/go-redis/redis"
)

func BuildGoRedis(config *Config) (*redislib.Client, error) {
	var (
		redisClient = redislib.NewClient(&redislib.Options{
			Addr: config.Addr,
		})
	)

	return redisClient, nil
}
