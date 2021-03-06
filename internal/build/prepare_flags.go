package build

import (
	"flag"
)

// PrepareFlags binds flags to configuration variables
func PrepareFlags(config *Config) {
	flag.StringVar(&config.Redis.Addr, "redis-addr", "redis:6379", "The address of the redis host")
	flag.StringVar(&config.Redis.Channel, "redis-channel", "global-room", "the redis channel to subscribe/publish to")
	flag.StringVar(&config.ServerAddr, "server-addr", "", "The address to listen on")
	flag.StringVar(&config.ServerMode, "server-mode", "ws", "The server mode (tcp, ws)")
	flag.StringVar(&config.JSONValidationSchemaPath, "json-validation-schema-path", "file:///tmp/message-schema.json", "The location of the JSON validation file")
	flag.StringVar(&config.Health.Addr, "health-addr", "", "The address to listen on for healthcheck requests")
	flag.StringVar(&config.Health.Path, "health-path", "/healthcheck", "The endpoint to response to for healthchecks")
}
