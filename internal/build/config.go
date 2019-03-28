package build

import "github.com/Kochava/k8s-demo-chat/internal/build/redis"

// Config stores application configuration variables
type Config struct {
	Redis                    *redis.Config
	ServerAddr               string
	JSONValidationSchemaPath string
}

// NewConfig initializes an empty config
func NewConfig() *Config {
	return &Config{
		Redis: &redis.Config{},
	}
}
