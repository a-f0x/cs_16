FROM golang:1.11.2-alpine3.8 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh

WORKDIR /bot
ADD ./ /bot

RUN cd /bot && \
    go get -d -v && \
    go build -v -o hlds_bot

FROM alpine:3.8

RUN apk update && \
  apk add ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/*

WORKDIR /bot
COPY --from=build-env /bot /bot

ENTRYPOINT exec ./hlds_bot
