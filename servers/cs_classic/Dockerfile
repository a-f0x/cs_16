#FROM golang:1.11.2-alpine3.8 AS build-env
#
#RUN apk update && apk upgrade && \
#    apk add --no-cache git openssh
#
#WORKDIR /game_runnner
#ADD ../hlds_starter /game_runnner
#
#RUN cd /discovery_service && \
#    go get -d -v && \
#    go build -v -o game_runnner
#

FROM af0x/hlds:latest
MAINTAINER Andrey Dzhigirey <af0xdev@gmail.com>
ARG RCON_PASSWORD

COPY mapcycle.txt /hlds/cstrike/mapcycle.txt
COPY server.cfg /hlds/cstrike/server.cfg
COPY ./start_hlds_server_game.sh /hlds/start_hlds_server_game.sh
COPY ./discovery_client /hlds/discovery_client
RUN echo rcon_password ${RCON_PASSWORD} >> /hlds/cstrike/server.cfg

WORKDIR /hlds

#ENTRYPOINT exec ./hlds_run -game cstrike +ip ${hlds_ip} +port ${hlds_port} +maxplayers 32 +map de_dust2 +logaddress ${log_receiver_ip} ${log_receiver_port}
#ENTRYPOINT exec ./game_runnner
ENTRYPOINT exec ./start_hlds_server_game.sh
