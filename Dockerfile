FROM golang:1.17

WORKDIR /usr/local/mafia-core

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY proto ./proto
COPY server ./server
COPY client ./client

RUN go build .