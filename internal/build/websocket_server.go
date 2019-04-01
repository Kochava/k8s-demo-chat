package build

import (
	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/websocketutil"
)

// WebsocketServer builds a Server based on the configuration
func WebsocketServer(config *Config) (*websocketutil.Server, error) {
	var (
		err              error
		readWriteHandler *broadcast.ReadWriteHandler
		websocketHandler websocketutil.Handler
	)

	if readWriteHandler, err = GetReadWriterHandler(config); err != nil {
		return nil, err
	}

	websocketHandler = &broadcast.WebsocketHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	return &websocketutil.Server{
		Addr:       config.ServerAddr,
		HandleFunc: websocketHandler.Handle,
	}, nil
}
