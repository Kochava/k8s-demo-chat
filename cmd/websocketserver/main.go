package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kochava/k8s-demo-chat/internal/build"
	"github.com/Kochava/k8s-demo-chat/internal/websocketutil"
)

func main() {
	var (
		err error

		config          = build.NewConfig()
		websocketServer *websocketutil.Server

		sigs = make(chan os.Signal, 1)
	)

	build.PrepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	if websocketServer, err = build.WebsocketServer(config); err != nil {
		log.Println("Unable to create websocket server:", err.Error())
		return
	}

	if err = websocketServer.ListenAndServe(); err != nil {
		log.Println("error starting server:", err.Error())
	}
}
