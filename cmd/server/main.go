// Command server starts a chat server
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/build"
	"github.com/Kochava/k8s-demo-chat/internal/build/healthcheck"
)

func main() {
	var (
		err error

		config            = build.NewConfig()
		broadcastServer   broadcast.Server
		healthcheckServer *http.Server

		sigs = make(chan os.Signal, 1)
	)

	build.PrepareFlags(config)
	flag.Parse()

	log.Println("Config", config)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	if broadcastServer, err = build.Server(config); err != nil {
		log.Println("Unable to create server:", err.Error())
		return
	}

	healthcheckServer = healthcheck.Build(config.Health)

	go startServer(broadcastServer)
	go startServer(healthcheckServer)

	// wait for a termination signal
	signal := <-sigs

	log.Printf("signal (%s) received, shutting down", signal)

	shutdownContext, shutdownContextCancelFunc := context.WithTimeout(context.Background(), time.Second)
	go healthcheckServer.Shutdown(shutdownContext)
	go broadcastServer.Shutdown(shutdownContext)

	// call the function to avoid linting error and to
	// reduce memory leaks
	defer shutdownContextCancelFunc()

	// wait for the shutdown duration
	<-shutdownContext.Done()
	log.Println("bye")
}
