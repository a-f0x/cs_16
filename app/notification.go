package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type TelegramNotificator struct {
	config *TelegramConfig
	bot    *tgbotapi.BotAPI
	chats  []*Chat
	//admins   []Chat
	selfName string
	BotEvent chan BotEvent
}

func NewTelegramNotificator(config *TelegramConfig) *TelegramNotificator {
	notificator := &TelegramNotificator{}
	notificator.config = config
	notificator.BotEvent = make(chan BotEvent)
	return notificator
}

func (n *TelegramNotificator) NotifyAll(message string) {
	for _, group := range n.chats {
		if !group.Muted {
			n.Notify(message, group.Id)
		}
	}
}

func (n *TelegramNotificator) Notify(message string, chatId int64) {
	n.sendMessage(message, chatId)
}

func (n *TelegramNotificator) tryConnect() *tgbotapi.BotAPI {
	log.Println("Trying to connect telegram bot api...")
	b, err := tgbotapi.NewBotAPIWithClient(n.config.Bot.Token, &http.Client{})
	if err != nil {
		log.Printf("Error connect to  telegram bot api: %s\nTry reconnect after %d sec...", err, n.config.Bot.ReconnectTimeout)
		time.Sleep(time.Duration(n.config.Bot.ReconnectTimeout) * time.Second)
		return n.tryConnect()
	}
	return b
}

func (n *TelegramNotificator) Start() {
	configureProxy(n.config.Proxy)
	n.initChats()
	n.bot = n.tryConnect()
	n.selfName = n.bot.Self.UserName
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	log.Printf("TelegramBot %s connected\n", n.selfName)

	updates, updateErr := n.bot.GetUpdatesChan(u)
	if updateErr != nil {
		log.Printf("Error receive update from  telegram bot: %s\n", updateErr)
		n.bot.StopReceivingUpdates()
		n.Start()

	} else {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
			n.processUpdate(update)
		}
	}

}

func (n *TelegramNotificator) processUpdate(update tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	groupName := update.Message.Chat.Title
	userName := update.Message.From.UserName
	text := update.Message.Text

	if update.Message.Chat.Type != "group" {
		n.processDirectMessage(chatId, userName, text)
		return
	}
	n.processGroupMessage(chatId, groupName, text)

}

func (n *TelegramNotificator) processDirectMessage(chatId int64, userName string, text string) {
	if !n.isChatExist(chatId) {
		n.addChat(userName, chatId, false, false)
		n.onAction(chatId, "", ChatAdded)
		return
	}
	if n.processBotCommands(text, chatId) {
		return
	}
	chat := n.getChat(chatId)
	if n.getChat(chatId).AllowExecuteRcon {
		n.onAction(chatId, text, RconCommand)
		return
	}

	if n.config.Bot.AdminPassword != text {
		n.sendMessage("Please enter the password", chatId)
		return
	}
	chat.AllowExecuteRcon = true
	n.serializeChats()
	n.sendMessage("You are added as the administrator. Write me a command and I will execute it on the server.\n"+
		"See list of commands http://cs1-6cfg.blogspot.com/p/cs-16-client-and-console-commands.html", chatId)
}

func (n *TelegramNotificator) processGroupMessage(chatId int64, groupName string, text string) {

	if !n.isChatExist(chatId) {
		n.addChat(groupName, chatId, false, false)
		n.onAction(chatId, "", ChatAdded)
		return
	}

	if !strings.Contains(text, n.selfName) {
		return
	}

	index := strings.Index(text, "@")
	message := text[:index]
	n.processBotCommands(message, chatId)
}

func (n *TelegramNotificator) processBotCommands(command string, chatId int64) bool {
	if command == "/mute" {
		n.muteChat(chatId, true)
		n.sendMessage("Chat muted", chatId)
		return true
	}

	if command == "/unmute" {
		n.muteChat(chatId, false)
		n.sendMessage("Chat unmuted", chatId)
		return true
	}

	if command == "/info" {
		n.onAction(chatId, command, BotCommand)
		return true
	}
	return false
}

func (n *TelegramNotificator) muteChat(chatId int64, mute bool) {
	chat := n.getChat(chatId)
	if chat == nil {
		return
	}
	chat.Muted = mute
	n.serializeChats()
}

func (n *TelegramNotificator) onAction(chatId int64, message string, action BotAction) {
	n.BotEvent <- BotEvent{
		ChatId:    chatId,
		BotAction: action,
		Message:   message,
	}
}

func configureProxy(proxyConfig TelegramProxyConfig) {
	if !proxyConfig.Enabled {
		return
	}
	proxyUrl := fmt.Sprintf("socks5://%s:%s@%s:%d",
		proxyConfig.User,
		proxyConfig.Password,
		proxyConfig.Host,
		proxyConfig.Port)

	err := os.Setenv("HTTP_PROXY", proxyUrl)
	if err != nil {
		panic(fmt.Errorf("Fatal error set os enb HTTP_PROXY: %s \n", err))
	}
}

func (n *TelegramNotificator) sendMessage(message string, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, message)
	_, e := n.bot.Send(msg)
	if e != nil {
		te, ok := e.(tgbotapi.Error)
		if ok {
			if te.Code == 403 || te.Code == 400 {
				n.removeChat(chatId)
				n.onAction(chatId, "", ChatRemoved)
			}
		}
		log.Println(fmt.Errorf("Error send message : %s \n", e))
	}
}

func (n *TelegramNotificator) initChats() {
	n.chats = make([]*Chat, 0)
	file, readError := ioutil.ReadFile("chats.json")
	if readError == nil {
		_ = json.Unmarshal(file, &n.chats)
	}
}

func (n *TelegramNotificator) addChat(chatName string, chatId int64, muted bool, allowRcon bool) {
	n.chats = append(n.chats, &Chat{chatName, chatId, muted, allowRcon})
	n.serializeChats()

}
func (n *TelegramNotificator) removeChat(chatId int64) {

	if !n.isChatExist(chatId) {
		return
	}

	chats := make([]*Chat, 0)
	for _, chat := range n.chats {
		if chat.Id == chatId {
			continue
		}
		chats = append(chats, chat)
	}
	n.chats = chats
	n.serializeChats()

}

func (n *TelegramNotificator) getChat(chatId int64) *Chat {
	for _, chat := range n.chats {
		if chat.Id == chatId {
			return chat
		}
	}
	return nil
}

func (n *TelegramNotificator) isChatExist(chatId int64) bool {
	return n.getChat(chatId) != nil
}

func (n *TelegramNotificator) serializeChats() {
	file, _ := json.MarshalIndent(n.chats, "", " ")
	_ = ioutil.WriteFile("chat_groups.json", file, 0644)
}
