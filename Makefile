PREFIX ?= $(HOME)
MAKE_PATH ?= $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

###############################################################################
# Update targets
###############################################################################
.PHONY: update
update: # Update submodules and other deps
	git submodule update --init --recursive
	go get -u golang.org/x/lint/golint

###############################################################################
# Server build targets
###############################################################################
servers: # Build the server binary
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o $(MAKE_PATH)bin/chat-server github.com/Kochava/k8s-demo-chat/cmd/server

###############################################################################
# Test targets
###############################################################################
.PHONY: test go-test go-lint
test: go-test go-lint

go-test:
	go test ./...

go-lint:
	golint ./{cmd,internal}/...

###############################################################################
# Cleaning targets
###############################################################################
.PHONY: clean
clean:
	rm -rf $(MAKE_PATH)bin/*

###############################################################################
# Local run targets
###############################################################################
.PHONY: composer-up
composer-up:
	docker-compose -f docker/docker-compose.yml --project-directory . up --build

.PHONY: composer-down
composer-down:
	docker-compose -f docker/docker-compose.yml --project-directory . down
