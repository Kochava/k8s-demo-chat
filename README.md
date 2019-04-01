# Demo Chat Server

[![CircleCI](https://circleci.com/gh/Kochava/k8s-demo-chat.svg?style=svg)](https://circleci.com/gh/Kochava/k8s-demo-chat)
[![Maintainability](https://api.codeclimate.com/v1/badges/fa5b1a1998ea2babbc0b/maintainability)](https://codeclimate.com/github/Kochava/k8s-demo-chat/maintainability)

The demo chat server is an applications that can support TCP or websocket connections, and a web server. The TCP and WebSocket servers publish and subscribe messages to Redis allowing the services to be stateless. The web server connects to the WebSocket server allowing an HTML interface and when running in TCP mode, the server can be connected to using any standard TCP based utility or software library.

## Usage

### Compile

```
make servers
```

### Running local

```
$ make composer-up
$ make composer-down
```
