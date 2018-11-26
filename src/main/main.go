package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const webhookURL = "https://grannys-bot.herokuapp.com/"

func main() {
	log.Println("Bot alives")
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("797530343:AAHNcti7m5XRKjTrDckANF7uXyoLZyTSxcA")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("Received text", update.Message.Text)

		message = tgbotapi.NewMessage(update.Message.Chat.ID, "hello")

		bot.Send(message)
	}
}
