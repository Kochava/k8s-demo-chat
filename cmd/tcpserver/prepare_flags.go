package main

import (
	"flag"

	"github.com/Kochava/k8s-demo-chat/internal/build"
)

func prepareFlags(config *build.Config) {
	flag.StringVar(&config.RedisAddr, "redis-addr", "redis:6379", "The address of the redis host")
	flag.StringVar(&config.RedisChannel, "redis-channel", "global-room", "the redis channel to subscribe/publish to")
	flag.StringVar(&config.ServerAddr, "server-addr", "", "The address to listen on")
	flag.StringVar(&config.JSONValidationSchemaPath, "json-validation-schema-path", "file:///tmp/message-schema.json", "The location of the JSON validation file")
}
