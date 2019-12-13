FROM golang:1.11.2-alpine3.8 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh

WORKDIR /app
ADD ./app /app

RUN cd /app && \
    go get -d -v && \
    go build -v -o cs_bot

FROM alpine:3.8

RUN apk update && \
  apk add ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=build-env /app/cs_bot /app

ENTRYPOINT exec ./cs_bot
