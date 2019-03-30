package main

import (
	"log"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
)

func startServer(server broadcast.Server) {
	if err := server.ListenAndServe(); err != nil {
		log.Println("server down:", err.Error())
	}
}
