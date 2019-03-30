package build

import (
	"github.com/Kochava/k8s-demo-chat/internal/build/healthcheck"
	"github.com/Kochava/k8s-demo-chat/internal/build/redis"
)

// Config stores application configuration variables
type Config struct {
	Redis                    *redis.Config
	ServerAddr               string
	ServerMode               string
	JSONValidationSchemaPath string
	Health                   *healthcheck.Config
}

// NewConfig initializes an empty config
func NewConfig() *Config {
	return &Config{
		Redis:  &redis.Config{},
		Health: &healthcheck.Config{},
	}
}
