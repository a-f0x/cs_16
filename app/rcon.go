package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

type Rcon struct {
	/*IP Address of HLDS server*/
	hldsServerAddress string
	/*Remote console password for HLDS server*/
	password []byte
}

func NewRcon(hldsServerAddress string, hldsServerPort int64, hldsPassword string) *Rcon {
	rcon := &Rcon{}
	rcon.hldsServerAddress = fmt.Sprintf("%s:%d", hldsServerAddress, hldsServerPort)
	rcon.password = []byte(hldsPassword)
	return rcon
}

func (client *Rcon) SendRconCommand(command string) (string, error) {

	response, err := client.send([]byte(command))
	if err != nil {
		return "", err
	}
	result := string(response[5:])
	if len(result) <= 2 {
		return "Wrong rcon command", nil
	}
	return result, nil

}

var (
	challenge = []byte("challenge rcon\n")
	prefix    = []byte("rcon")
	header    = []byte{0xFF, 0xFF, 0xFF, 0xFF}
	delimiter = []byte{32}
)

func (client *Rcon) send(data []byte) ([]byte, error) {

	connection, err := client.openConnection()
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	commandId, err := commandId(connection)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.Write(prefix)
	buf.Write(delimiter)
	buf.Write(commandId)
	buf.Write(delimiter)
	buf.Write(client.password)
	buf.Write(delimiter)
	buf.Write(data)

	return write(connection, buf.Bytes())

}

func (client *Rcon) openConnection() (*net.UDPConn, error) {
	server, err := net.ResolveUDPAddr("udp4", client.hldsServerAddress)
	if err != nil {
		return nil, err
	}
	connection, err := net.DialUDP("udp4", nil, server)
	if err != nil {
		return nil, err
	}
	deadline := time.Now().Add(time.Duration(1) * time.Second)
	connection.SetDeadline(deadline)
	return connection, nil

}

func commandId(connection *net.UDPConn) ([]byte, error) {
	response, err := write(connection, challenge)
	if err != nil {
		return nil, err
	}
	i := bytes.LastIndexByte(response, 32)

	return response[i+1 : len(response)-1], nil
}

func write(connection *net.UDPConn, data []byte) ([]byte, error) {

	buf := new(bytes.Buffer)

	buf.Write(header)
	buf.Write(data)

	_, err := connection.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 2048)

	n, _, err := connection.ReadFrom(buffer)

	if err != nil {
		return nil, err
	}
	return buffer[0 : n-1], nil
}
