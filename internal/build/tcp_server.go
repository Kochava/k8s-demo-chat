package build

import (
	"github.com/Kochava/k8s-demo-chat/internal/broadcast"
	"github.com/Kochava/k8s-demo-chat/internal/tcputil"
)

func TCPServer(config *Config) (*tcputil.Server, error) {
	var (
		err error

		readWriteHandler *broadcast.ReadWriteHandler
		tcpHandler       tcputil.Handler
	)

	if readWriteHandler, err = GetReadWriterHandler(config); err != nil {
		return nil, err
	}

	tcpHandler = &broadcast.TCPHandlerProxy{
		ReadWriteHandler: readWriteHandler,
	}

	return &tcputil.Server{
		Addr:    config.ServerAddr,
		Handler: tcpHandler,
	}, nil
}
