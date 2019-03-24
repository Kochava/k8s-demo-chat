FROM alpine:3.9.2
COPY bin/chat-tcp-server /usr/local/bin/chat-tcp-server

EXPOSE 4000

CMD ["/tmp/bin/chat-tcp-server", "-server-addr", "0.0.0.0:4000"]