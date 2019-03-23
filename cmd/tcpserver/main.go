package main

import (
	"flag"
	"log"

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
	)

	prepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	if readWriteHandler, err = build.GetReadWriterHandler(config); err != nil {
		log.Println("unable to get read write handler:", err.Error())
		return
	}

	tcpHandler = &broadcast.TCPHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	tcpServer = tcputil.Server{
		Addr:    config.ServerAddr,
		Handler: tcpHandler,
	}

	if err = tcpServer.ListenAndServe(); err != nil {
		log.Println("error starting server:", err.Error())
	}
}
