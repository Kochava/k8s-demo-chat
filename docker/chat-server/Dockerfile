FROM alpine:3.9.2
RUN mkdir -p /etc/chat
COPY bin/chat-server /usr/local/bin/chat-server
COPY contrib/message-schema.json /etc/chat/message-schema.json

CMD ["/usr/local/bin/chat-server"]
