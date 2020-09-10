#!/usr/bin/env sh
./discovery_client
./hlds_run -game cstrike +ip ${hlds_ip} +port ${hlds_port} +maxplayers 32 +map de_dust2 +logaddress ${log_receiver_ip} ${log_receiver_port}