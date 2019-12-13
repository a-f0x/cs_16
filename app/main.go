package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	notificator                    *TelegramNotificator
	rcon                           *Rcon
	csIp                           string
	csPort                         int64
	logReceiverPort                int64
	rconPassword                   string
	isTelegramProxyEnabled         bool
	telegramProxyUser              string
	telegramProxyPassword          string
	telegramProxyHost              string
	telegramProxyPort              int64
	telegramBotToken               string
	telegramBotReconnectTimeoutSec int64
	telegramAdminPassword          string
)

func main() {

	csIp = os.Getenv("cs_ip")
	var err error = nil

	csPort, err = strconv.ParseInt(os.Getenv("cs_port"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("cs_port %s not valid \n", os.Getenv("cs_port")))
	}

	logReceiverPort, err = strconv.ParseInt(os.Getenv("log_receiver_port"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("log_receiver_port %s not valid \n", os.Getenv("log_receiver_port")))
	}

	isTelegramProxyEnabled, err = strconv.ParseBool(os.Getenv("proxy_enabled"))
	if err != nil {
		panic(fmt.Errorf("proxy_enabled %s not valid \n", os.Getenv("proxy_enabled")))
	}
	telegramProxyHost = os.Getenv("proxy_host")

	telegramProxyPort, err = strconv.ParseInt(os.Getenv("proxy_port"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("proxy_port %s not valid \n", os.Getenv("proxy_port")))
	}

	telegramProxyUser = os.Getenv("proxy_user")

	telegramProxyPassword = os.Getenv("proxy_password")

	telegramBotToken = os.Getenv("bot_token")

	telegramBotReconnectTimeoutSec, err = strconv.ParseInt(os.Getenv("bot_reconnect_timeout"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("bot_reconnect_timeout %s not valid \n", os.Getenv("bot_reconnect_timeout")))
	}

	telegramAdminPassword = os.Getenv("admin_password")

	rconPassword = os.Getenv("rcon_password")

	rcon = NewRcon(csIp, csPort, rconPassword)

	initNotificator(&TelegramConfig{

		Proxy: TelegramProxyConfig{

			isTelegramProxyEnabled,
			telegramProxyHost,
			telegramProxyPort,
			telegramProxyUser,
			telegramProxyPassword},

		Bot: TelegramBotConfig{
			telegramBotToken,
			telegramBotReconnectTimeoutSec,
			telegramAdminPassword,
		},
	})

	initLogReceiver(logReceiverPort)
}

func initNotificator(config *TelegramConfig) {

	notificator = NewTelegramNotificator(config)
	go func() {
		for {
			select {
			case event := <-notificator.BotEvent:
				switch action := event.BotAction; action {
				case ChatAdded:
					notificator.Notify("Chat added", event.ChatId)
					//i, err := getFullServerInfo()
					//if err != nil {
					//	notificator.Notify("Error get server info:"+err.Error(), event.ChatId)
					//} else {
					//	notificator.Notify(i, event.ChatId)
					//}

				case BotCommand:
					result, err := getFullServerInfo()
					if err != nil {
						notificator.Notify(err.Error(), event.ChatId)
					} else {
						notificator.Notify(result, event.ChatId)
					}
				case RconCommand:
					result, err := rcon.SendRconCommand(event.Message)
					if err != nil {
						notificator.Notify(err.Error(), event.ChatId)
					} else {
						notificator.Notify(result, event.ChatId)
					}

				}
			}
		}
	}()

	go func() {
		notificator.Start()
	}()
}

func getFullServerInfo() (string, error) {

	return rcon.SendRconCommand("status")
}

func initLogReceiver(port int64) {

	logReceiver := NewLogReceiver(port)
	go func() {
		for {
			select {
			case event := <-logReceiver.GameEvent:
				notificator.NotifyAll(event.String())
			}
		}
	}()

	logReceiver.Start()
}
