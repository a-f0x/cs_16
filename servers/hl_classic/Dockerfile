FROM af0x/hlds:latest
MAINTAINER Andrey Dzhigirey <af0xdev@gmail.com>
ARG RCON_PASSWORD

COPY ./server.cfg /hlds/valve/server.cfg
COPY ./mapcycle.txt /hlds/valve/mapcycle.txt
COPY ./start_hlds_server_game.sh /hlds/start_hlds_server_game.sh
COPY ./discovery_client /hlds/discovery_client
RUN echo rcon_password ${RCON_PASSWORD} >> /hlds/valve/server.cfg

WORKDIR /hlds
ENTRYPOINT exec ./start_hlds_server_game.sh
#ENTRYPOINT exec ./hlds_run -game valve +ip ${hlds_ip} +port ${hlds_port} +maxplayers 16 +map crossfire +exec server.cfg +logaddress ${log_receiver_ip} ${log_receiver_port}
