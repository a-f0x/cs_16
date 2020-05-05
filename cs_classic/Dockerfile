FROM af0x/hlds:latest
MAINTAINER Andrey Dzhigirey <af0xdev@gmail.com>
ARG RCON_PASSWORD

COPY ./cstrike/mapcycle.txt /hlds/cstrike/mapcycle.txt
COPY ./cstrike/server.cfg /hlds/cstrike/server.cfg
RUN echo rcon_password ${RCON_PASSWORD} >> /hlds/cstrike/server.cfg
WORKDIR /hlds

ENTRYPOINT exec ./hlds_run -game cstrike +ip ${hlds_ip} +port ${hlds_port} +maxplayers 32 +map de_dust2 +logaddress ${log_receiver_ip} ${log_receiver_port}
