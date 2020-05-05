package main

type GameRunner struct {
	hLDSServer           HLDSServer
	discoveryServiceIp   string
	discoveryServicePort string
}

func NewGameRunner(
	hldsServer HLDSServer,
	discoveryServiceIp string,
	discoveryServicePort string) *GameRunner {

	return &GameRunner{
		hLDSServer:           hldsServer,
		discoveryServiceIp:   discoveryServiceIp,
		discoveryServicePort: discoveryServicePort,
	}
}
