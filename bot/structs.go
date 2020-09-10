package main

import (
	"fmt"
	"time"
)

type TelegramProxyConfig struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type TelegramBotConfig struct {
	Token            string `json:"token"`
	ReconnectTimeout int64  `json:"reconnect_timeout"`
	AdminPassword    string `json:"admin_password"`
}

type TelegramConfig struct {
	Proxy TelegramProxyConfig `json:"proxy"`
	Bot   TelegramBotConfig   `json:"bot"`
}

type Chat struct {
	Name             string
	Id               int64
	Muted            bool
	AllowExecuteRcon bool
}

type GameEvent struct {
	TimeStamp    time.Time
	PlayerAction PlayerAction
	Nick         string
}

func (e *GameEvent) String() string {
	return fmt.Sprintf("Player %s %s.", e.Nick, ops[e.PlayerAction])
}

type PlayerAction uint32

const (
	Connected    PlayerAction = 0
	Disconnected PlayerAction = 1
)

var ops = map[PlayerAction]string{
	Connected:    "Connected",
	Disconnected: "Disconnected",
}

type BotEvent struct {
	ChatId    int64
	BotAction BotAction
	Message   string
}

type BotAction uint32

const (
	ChatAdded   BotAction = 0
	ChatRemoved BotAction = 1
	RconCommand BotAction = 3
)

type HLDSServer struct {
	Name            string `json:"name"`
	Game            string `json:"game"`
	RconPassword    string `json:"rcon_password"`
	HLDSIp          string `json:"hlds_ip"`
	HLDSPort        string `json:"hlds_port"`
	LogReceiverPort string `json:"log_receiver_port"`
}

type HLDSServers []HLDSServer

type DiscoveryEvent struct {
	HLDSServer HLDSServer
}
