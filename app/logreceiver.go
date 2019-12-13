package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type LogReceiver struct {
	port      int64
	GameEvent chan GameEvent
}

func NewLogReceiver(port int64) LogReceiver {
	return LogReceiver{
		port:      port,
		GameEvent: make(chan GameEvent),
	}
}

func (l *LogReceiver) Start() {
	udpServer, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s%d", ":", l.port))

	if err != nil {
		log.Fatalf("Error start udp server for listen event from cs server: %s \n", err)
	}

	connection, err := net.ListenUDP("udp4", udpServer)

	defer connection.Close()

	if err != nil {
		log.Fatalf("Error start udp server for listet event from cs server: %s \n", err)
	}

	buffer := make([]byte, 1024)
	log.Println("Log receiver started")

	for {
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("Error read bytes from server: %s \n", err)
		}
		l.parseLogEvent(buffer[10 : n-1])
	}
}

func (l *LogReceiver) parseLogEvent(event []byte) {
	if len(event) < 24 {
		return
	}
	//log.Println(string(event))
	str := string(event[23:])
	if strings.Contains(str, " entered the game") {
		l.GameEvent <- GameEvent{
			TimeStamp:    time.Now(),
			PlayerAction: Connected,
			Nick:         findNick(str),
		}

	}

	if strings.Contains(str, " disconnected") {
		l.GameEvent <- GameEvent{
			TimeStamp:    time.Now(),
			PlayerAction: Disconnected,
			Nick:         findNick(str),
		}
	}
}

func findNick(eventString string) string {
	index := strings.Index(eventString, "<")
	return eventString[:index]

}
