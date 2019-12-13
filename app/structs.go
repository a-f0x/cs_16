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
	BotCommand  BotAction = 2
	RconCommand BotAction = 3
)
