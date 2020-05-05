FROM af0x/halflife:latest
MAINTAINER Andrey Dzhigirey <af0xdev@gmail.com>
ARG RCON_PASSWORD

COPY ./halflife/server.cfg /hl/valve/server.cfg
RUN echo rcon_password ${RCON_PASSWORD} >> /hl/valve/server.cfg
WORKDIR /hl

ENTRYPOINT exec ./hlds_run -game valve +ip ${hlds_ip} +port ${hlds_port} +maxplayers 16 +map crossfire +exec server.cfg +logaddress ${log_receiver_ip} ${log_receiver_port}
