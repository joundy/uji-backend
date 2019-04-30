FROM golang:1.8

WORKDIR /go/src/github.com/haffjjj/uji-backend/
COPY . .

CMD ["make", "start"]