FROM golang:1.16.13-alpine3.15

RUN mkdir /app
WORKDIR /app

RUN go get github.com/cosmtrek/air@v1.29.0

ADD ./gin-clean-architecture-format/go.mod /app
RUN go mod download
