PREFIX ?= $(HOME)
MAKE_PATH ?= $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: servers
servers: servers-websockets servers-tcp

.PHONY: servers-websockets
servers-websockets:
	GOOS=linux GOARCH=amd64 go build -o $(MAKE_PATH)bin/chat-ws-server github.com/Kochava/k8s-demo-chat/cmd/websocketserver

.PHONY: servers-tcp
servers-tcp:
	GOOS=linux GOARCH=amd64 go build -o $(MAKE_PATH)bin/chat-tcp-server github.com/Kochava/k8s-demo-chat/cmd/tcpserver

.PHONY: clean
clean:
	rm -rf $(MAKE_PATH)bin/*

.PHONY: run-local
run:
	# docker compose here
