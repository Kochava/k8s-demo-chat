package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/build"
)

func main() {
	var (
		err error

		config = build.NewConfig()
		server broadcast.Server

		sigs = make(chan os.Signal, 1)
	)

	build.PrepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	if server, err = build.Server(config); err != nil {
		log.Println("Unable to create server:", err.Error())
		return
	}

	if err = server.ListenAndServe(); err != nil {
		log.Println("error starting server:", err.Error())
	}
}
