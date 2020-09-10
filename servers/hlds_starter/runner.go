package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DiscoveryClient struct {
	hLDSServer HLDSServer
	serverUrl  string
}

func NewDiscoveryClient(
	hldsServer HLDSServer,
	discoveryServiceIp string,
	discoveryServicePort string) *DiscoveryClient {
	url := "http://" + discoveryServiceIp + ":" + discoveryServicePort + "/"
	return &DiscoveryClient{
		hLDSServer: hldsServer,
		serverUrl:  url,
	}
}

func (gr *DiscoveryClient) RegisterGameOnServer() {
	if gr.registerGame() {
		fmt.Printf("Game is successfully registered on discovery server %s \n", gr.serverUrl)
	}
}

func (gr *DiscoveryClient) registerGame() bool {
	log.Printf("Trying to register game on %s...", gr.serverUrl)

	jsonStr, err := json.Marshal(gr.hLDSServer)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", gr.serverUrl+"servers", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error connect to  discovery service api: %s\nTry reconnect after 3 sec...", err)

		time.Sleep(time.Duration(3) * time.Second)
		return gr.registerGame()
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return true

}
