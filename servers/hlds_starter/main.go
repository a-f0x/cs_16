package main

import "os"

func fakeEnv() {
	os.Setenv("hlds_ip", "111.111.111.1")
	os.Setenv("hlds_port", "123")
	os.Setenv("log_receiver_port", "22")
	os.Setenv("log_receiver_ip", "222.222.222.2")
	os.Setenv("game", "fake")
	os.Setenv("name", "fake")
	os.Setenv("discovery_ip", "127.0.0.1")
	os.Setenv("discovery_port", "8030")
	os.Setenv("rcon_password", "pwd")

}
func main() {
	fakeEnv()
	var hldsIp = os.Getenv("hlds_ip")
	var hldsPort = os.Getenv("hlds_port")
	var logReceiverPort = os.Getenv("log_receiver_port")
	var logReceiverIp = os.Getenv("log_receiver_ip")
	var game = os.Getenv("game")
	var name = os.Getenv("name")
	var discoveryServiceHost = os.Getenv("discovery_ip")
	var discoveryServicePort = os.Getenv("discovery_port")
	var rconPassword = os.Getenv("rcon_password")

	server := HLDSServer{
		Name:            name,
		Game:            game,
		RconPassword:    rconPassword,
		HLDSIp:          hldsIp,
		HLDSPort:        hldsPort,
		LogReceiverIp:   logReceiverIp,
		LogReceiverPort: logReceiverPort,
	}

	runner := NewDiscoveryClient(server, discoveryServiceHost, discoveryServicePort)

	runner.RegisterGameOnServer()
}
