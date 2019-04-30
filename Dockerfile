FROM golang:1.8

COPY . /go/src/github.com/haffjjj/uji-backend
WORKDIR /go/src/github.com/haffjjj/uji-backend

CMD make start

EXPOSE 9001