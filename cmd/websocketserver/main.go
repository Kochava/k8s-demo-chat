package main

import (
	"flag"
	"log"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/build"
	"github.com/Kochava/k8s-demo-chat/internal/websocketutil"
)

func main() {
	var (
		err error

		config = &build.Config{}

		readWriteHandler *broadcast.ReadWriteHandler
		websocketHandler websocketutil.Handler

		websocketServer websocketutil.Server
	)

	prepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	if readWriteHandler, err = build.GetReadWriterHandler(config); err != nil {
		log.Println("unable to get read write handler:", err.Error())
		return
	}

	websocketHandler = &broadcast.WebsocketHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	websocketServer = websocketutil.Server{
		Addr:       config.ServerAddr,
		HandleFunc: websocketHandler.Handle,
	}

	if err = websocketServer.ListenAndServe(); err != nil {
		log.Println("error starting server:", err.Error())
	}
}
