package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var port string
	port = os.Getenv("port")
	if len(port) == 0 {
		port = "8030"
	}
	log.Printf("Discovery service starts at port %s \n", port)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
