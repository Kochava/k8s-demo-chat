package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

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

		sigs = make(chan os.Signal, 1)
	)

	prepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	if readWriteHandler, err = build.GetReadWriterHandler(config); err != nil {
		log.Println("unable to get read write handler:", err.Error())
		return
	}

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	websocketHandler = &broadcast.WebsocketHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	websocketServer = websocketutil.Server{
		Addr:       config.ServerAddr,
		HandleFunc: websocketHandler.Handle,
	}

	go func() {
		if err = websocketServer.ListenAndServe(); err != nil {
			log.Fatalf("error starting server: %s", err.Error())
		}
	}()

	<-sigs
	websocketServer.Shutdown(context.Background())
}
