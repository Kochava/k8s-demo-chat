package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kochava/k8s-demo-chat/internal/build"
	"github.com/Kochava/k8s-demo-chat/internal/tcputil"
)

func main() {
	var (
		err error

		config    = build.NewConfig()
		tcpServer *tcputil.Server

		sigs = make(chan os.Signal, 1)
	)

	build.PrepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	if tcpServer, err = build.TCPServer(config); err != nil {
		log.Println("unable to make TCP server:", err.Error())
		return
	}

	if err = tcpServer.ListenAndServe(); err != nil {
		log.Println("error starting server:", err.Error())
	}
}
