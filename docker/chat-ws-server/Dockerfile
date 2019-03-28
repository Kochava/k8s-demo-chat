FROM alpine:3.9.2
COPY bin/chat-ws-server /usr/local/bin/chat-ws-server

EXPOSE 4001

CMD ["/usr/local/bin/chat-ws-server", "-server-addr", "0.0.0.0:4001"]