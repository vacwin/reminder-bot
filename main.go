package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// get token from @BotFather
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is empty")
	}
	// setup bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("authorized as @%s", bot.Self.UserName)
	// get updates (?)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	//
	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := strings.TrimSpace(update.Message.Text)
		if text == "" {
			continue
		}
		switch {
		case strings.HasPrefix(text, "/start"):
			send(
				bot,
				update.Message.Chat.ID,
				"start",
			)
		case strings.HasPrefix(text, "/help"):
			send(
				bot,
				update.Message.Chat.ID,
				"help",
			)
		}
	}
}

func send(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = bot.Send(msg)
}
