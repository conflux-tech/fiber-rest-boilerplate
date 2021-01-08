FROM golang:1.14-alpine

RUN apk update && apk add curl \
build-base \
git \
make

WORKDIR /app

COPY ./ /app

RUN go build -o ./bin/fiber-rest-boilerplate ./
