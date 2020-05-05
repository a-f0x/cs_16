package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"RegisterHLDSServer",
		"POST",
		"/servers",
		RegisterHLDSServer,
	},
	Route{
		"ListRegisteredHLDSServers",
		"GET",
		"/servers",
		ListRegisteredHLDSServers,
	},
	Route{
		Name:        "RegisterCallBack",
		Method:      "POST",
		Pattern:     "/callbacks",
		HandlerFunc: RegisterCallBack,
	},
}
