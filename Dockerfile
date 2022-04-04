# syntax=docker/dockerfile:1
FROM golang:1.17.8-alpine

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o /api

EXPOSE 8000

CMD ["/api"]
