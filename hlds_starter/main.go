package main

import "os"

var server HLDSServer

func main() {
	var hldsIp = os.Getenv("hlds_ip")
	var hldsPort = os.Getenv("hlds_port")
	var logReceiverPort = os.Getenv("log_receiver_port")
	var logReceiverIp = os.Getenv("log_receiver_ip")
	var game = os.Getenv("game")
	var name = os.Getenv("name")
	var discoveryServiceHost = os.Getenv("discovery_host")
	var discoveryServicePort = os.Getenv("discovery_port")
	var rconPassword = os.Getenv("rcon_password")

	server = HLDSServer{
		Name:            name,
		Game:            game,
		RconPassword:    rconPassword,
		HLDSIp:          hldsIp,
		HLDSPort:        hldsPort,
		LogReceiverIp:   logReceiverIp,
		LogReceiverPort: logReceiverPort,
	}

}
