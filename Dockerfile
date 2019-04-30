# Builder
FROM golang:1.11.4-alpine3.8

RUN apk update && apk upgrade && \
    apk --update add git gcc make

WORKDIR /go/src/github.com/haffjjj/uji-backend

COPY . .

CMD make start