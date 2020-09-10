#!/usr/bin/env sh
./discovery_client
./hlds_run -game valve +ip ${hlds_ip} +port ${hlds_port} +maxplayers 16 +map crossfire +exec server.cfg +logaddress ${log_receiver_ip} ${log_receiver_port}