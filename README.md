# Distchat

Distchat contains three applications: a TCP server, a WebSocket server, and
a web server. The TCP and WebSocket servers publish and subscribe messages to
Redis allowing the services to be stateless. The web server connects to the
WebSocket server allowing an HTML interface. The TCP server can be connected to
using `nc` or similar applications.

## Usage

### Start

```
$ bash ./compile.sh
$ docker-compose up
```
### Chat

**TCP**

```
$ nc 127.0.0.1 4000
```

**Web**

http://127.0.0.1:8080/
