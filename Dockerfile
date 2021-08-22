FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/informer
COPY . /go/src/informer
WORKDIR /go/src/informer

RUN go get

CMD go run main.go