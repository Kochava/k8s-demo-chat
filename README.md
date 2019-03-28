# Distchat

[![CircleCI](https://circleci.com/gh/Kochava/k8s-demo-chat.svg?style=svg)](https://circleci.com/gh/Kochava/k8s-demo-chat)

Distchat contains three applications: a TCP server, a WebSocket server, and
a web server. The TCP and WebSocket servers publish and subscribe messages to
Redis allowing the services to be stateless. The web server connects to the
WebSocket server allowing an HTML interface. The TCP server can be connected to
using `nc` or similar applications.

## Usage

### Compile

```
make servers
```

### Container Image Creation

```
docker build -f docker/Dockerfile.frontend -t k8sdemo-frontend .
docker build -f docker/Dockerfile.chat-tcp-server -t k8sdemo-chat-tcp-server .
docker build -f docker/Dockerfile.chat-ws-server -t k8sdemo-chat-ws-server .

docker tag k8sdemo-frontend gcr.io/<project>/k8sdemo-frontend/latest
docker tag k8sdemo-chat-tcp-server gcr.io/<project>/k8sdemo-chat-tcp-server/latest
docker tag k8sdemo-chat-ws-server gcr.io/<project>/k8sdemo-chat-ws-server/latest

docker push gcr.io/<project>/k8sdemo-frontend/latest
docker push gcr.io/<project>/k8sdemo-chat-tcp-server/latest
docker push gcr.io/<project>/k8sdemo-chat-ws-server/latest
```

### Running local

```
$ make run
```

### Chat

**TCP**

```
$ nc 127.0.0.1 4000
```

**Web**

http://127.0.0.1:8080/

