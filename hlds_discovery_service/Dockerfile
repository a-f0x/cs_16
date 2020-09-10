FROM golang:1.11.2-alpine3.8 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh

WORKDIR /hlds_discovery_service
ADD ./ /hlds_discovery_service

RUN cd /hlds_discovery_service && \
    go get -d -v && \
    go build -v -o hlds_discovery_service

FROM alpine:3.8

RUN apk update && \
  apk add ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/*

WORKDIR /hlds_discovery_service
COPY --from=build-env /hlds_discovery_service /hlds_discovery_service

ENTRYPOINT exec ./hlds_discovery_service
