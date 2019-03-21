package main

import (
	"flag"

	"git.dev.kochava.com/notter/distchat/build"
)

func prepareFlags(config *build.Config) {
	flag.StringVar(&config.RedisAddr, "redis-addr", "redis:6379", "The address of the redis host")
	flag.StringVar(&config.RedisChannel, "redis-channel", "global-room", "the redis channel to subscribe/publish to")
	flag.StringVar(&config.ServerAddr, "server-addr", "", "The address to listen on")
}
