package main

type HLDSServer struct {
	Name            string `json:"name"`
	Game            string `json:"game"`
	RconPassword    string `json:"rcon_password"`
	HLDSIp          string `json:"hlds_ip"`
	HLDSPort        string `json:"hlds_port"`
	LogReceiverIp   string `json:"log_receiver_ip"`
	LogReceiverPort string `json:"log_receiver_port"`
}

type HLDSServers []HLDSServer
