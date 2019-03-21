#!/bin/bash

# ls -alh /go/src/git.dev.kochava.com/notter/distchat/vendor

GOOS=linux GOARCH=amd64 go build -o ./bin/tcpserver git.dev.kochava.com/notter/distchat/cmd/tcpserver
GOOS=linux GOARCH=amd64 go build -o ./bin/websocketserver git.dev.kochava.com/notter/distchat/cmd/websocketserver
