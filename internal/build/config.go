package build

import "github.com/Kochava/k8s-demo-chat/internal/build/redis"

// Config stores application configuration variables
type Config struct {
	Redis                    *redis.Config
	ServerAddr               string
	JSONValidationSchemaPath string
}
