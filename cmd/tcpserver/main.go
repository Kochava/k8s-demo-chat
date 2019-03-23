package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/build"
	"github.com/Kochava/k8s-demo-chat/internal/tcputil"
)

func main() {
	var (
		err error

		config = &build.Config{}

		readWriteHandler *broadcast.ReadWriteHandler
		tcpHandler       tcputil.Handler

		tcpServer tcputil.Server

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

	tcpHandler = &broadcast.TCPHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	tcpServer = tcputil.Server{
		Addr:    config.ServerAddr,
		Handler: tcpHandler,
	}

	go func() {
		if err = tcpServer.ListenAndServe(); err != nil {
			log.Fatalf("error starting server: %s", err.Error())
		}
	}()

	<-sigs
	tcpServer.Shutdown()
}
