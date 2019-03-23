PREFIX ?= $(HOME)
MAKE_PATH ?= $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: servers # Build server binaries
servers: servers-websockets servers-tcp

.PHONY: servers-websockets
servers-websockets: # Build the websockets binary
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o $(MAKE_PATH)bin/chat-ws-server github.com/Kochava/k8s-demo-chat/cmd/websocketserver

.PHONY: servers-tcp
servers-tcp: # Build the tcp server binary
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o $(MAKE_PATH)bin/chat-tcp-server github.com/Kochava/k8s-demo-chat/cmd/tcpserver

.PHONY: update
update: # Update submodules and other deps
	git submodule update --init --recursive
	go get -u golang.org/x/lint/golint

.PHONY: test
test: go-test go-lint

.PHONY: go-test
go-test:
	go test ./...

.PHONY: go-lint
go-lint:
	golint ./{cmd,internal}/...

.PHONY: clean
clean:
	rm -rf $(MAKE_PATH)bin/*

.PHONY: run-local
run:
	docker-compose -f docker/docker-compose.yml --project-directory . up

stop:
	docker-compose -f docker/docker-compose.yml --project-directory . down
