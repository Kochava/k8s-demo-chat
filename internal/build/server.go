package build

import (
	"errors"

	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
)

// Server builds a broadcast server based on the configuration
func Server(config *Config) (broadcast.Server, error) {
	switch config.ServerMode {
	case "tcp":
		return TCPServer(config)
	case "ws":
		return WebsocketServer(config)
	}

	return nil, errors.New("unsupported server mode")
}
