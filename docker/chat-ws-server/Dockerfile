FROM alpine:3.9.2
COPY bin/chat-tcp-server /usr/local/bin/chat-tcp-server

EXPOSE 4001

CMD ["/tmp/bin/chat-ws-server", "-server-addr", "0.0.0.0:4001"]