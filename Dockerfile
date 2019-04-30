FROM golang:1.8

ENV GOBIN /go/bin
ENV GOPATH /go

WORKDIR /go/src/github.com/haffjjj/uji-backend/
COPY . .

CMD ["make", "start"]