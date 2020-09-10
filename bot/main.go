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

func fakeEnvs() {
	os.Setenv("hlds_ip", "192.168.88.201")
	os.Setenv("hlds_port", "27015")
	os.Setenv("log_receiver_port", "10051")
	os.Setenv("proxy_enabled", "true")
	os.Setenv("proxy_host", "31.14.133.176")
	os.Setenv("proxy_port", "61555")
	os.Setenv("proxy_user", "user")
	os.Setenv("proxy_password", "rR4oa6ZUjIo")
	os.Setenv("bot_token", "412092394:AAF1ZsrasnOYStLor7nsBf8D07nal11ClaI")
	os.Setenv("bot_reconnect_timeout", "15")
	os.Setenv("admin_password", "admin_f0x")
	os.Setenv("rcon_password", "asjop2340239857uG")

}

func main() {
	fakeEnvs()
	csIp = os.Getenv("hlds_ip")
	var err error = nil

	csPort, err = strconv.ParseInt(os.Getenv("hlds_port"), 10, 64)
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
