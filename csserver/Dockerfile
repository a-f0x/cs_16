FROM af0x/cs_server:csdm
MAINTAINER Andrey Dzhigirey <af0xdev@gmail.com>
ARG RCON_PASSWORD

COPY ./cstrike/addons/amxmodx /opt/cs16-server/cstrike/addons/amxmodx
COPY ./cstrike/maps /opt/cs16-server/cstrike/maps
COPY ./cstrike/mapcycle.txt /opt/cs16-server/cstrike/mapcycle.txt
COPY ./cstrike/server.cfg /opt/cs16-server/cstrike/server.cfg
RUN echo rcon_password ${RCON_PASSWORD} >> /opt/cs16-server/cstrike/server.cfg
WORKDIR /opt/cs16-server

ENTRYPOINT exec ./hlds_run -game cstrike +ip ${cs_ip} +port ${cs_port} +maxplayers 32 +map cs_16fighter +logaddress ${log_receiver_ip} ${log_receiver_port}
