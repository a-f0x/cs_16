FROM golang:1.11.2-alpine3.8 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh

WORKDIR /discovery_service
ADD ./ /hlds_discovery_service

RUN cd /discovery_service && \
    go get -d -v && \
    go build -v -o hlds_discovery_service

FROM alpine:3.8

RUN apk update && \
  apk add ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/cache/apk/*

WORKDIR /discovery_service
COPY --from=build-env /discovery_service /discovery_service

ENTRYPOINT exec ./hlds_discovery_service
