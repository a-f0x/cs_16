package main

var servers HLDSServers
var callBacks Callbacks

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

func RemoveServer(name string) {

}

func CreateCallBack(callback Callback) Callback {
	for _, value := range callBacks {
		if value.Url == callback.Url {
			return callback
		}
	}
	callBacks = append(callBacks, callback)
	return callback
}
