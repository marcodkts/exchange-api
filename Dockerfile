# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /go/src/exchange-api

COPY . /go/src/exchange-api

# RUN export GOROOT=/go
# RUN go get exchange-api
# RUN go install


EXPOSE 8080