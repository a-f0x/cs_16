package main

var servers HLDSServers

func CreateHLDSServer(s HLDSServer) HLDSServer {
	for _, value := range servers {
		if value.Name == s.Name {
			return s
		}
	}
	servers = append(servers, s)
	return s
}

func GetAllHLDSServers() []HLDSServer {
	values := make([]HLDSServer, 0, len(servers))
	for _, v := range servers {
		values = append(values, v)
	}
	return values
}
