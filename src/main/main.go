package main

import (
	"fmt"
	"log"
	"net/http"
)

const token = "797530343:AAHNcti7m5XRKjTrDckANF7uXyoLZyTSxcA"
const webhookURL = "https://grannys-bot.herokuapp.com/"
const url = "https://api.telegram.org/"

func main() {
	var resp *http.Response
	resp, err := http.Get(url + token + "/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Body)

	// log.Println("Bot alives")
	// port := os.Getenv("PORT")
	// bot, err := tgbotapi.NewBotAPI("797530343:AAHNcti7m5XRKjTrDckANF7uXyoLZyTSxcA")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bot.Debug = true

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	// _, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// updates := bot.ListenForWebhook("/")
	// go http.ListenAndServe(":"+port, nil)

	// for update := range updates {

	// 	var message tgbotapi.MessageConfig
	// 	log.Println("Received text", update.Message.Text)

	// 	message = tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
	// 	bot.Send(message)

	// 	var pin tgbotapi.PinChatMessageConfig
	// 	pin.ChatID = update.Message.Chat.ID
	// 	pin.MessageID = update.Message.MessageID
	// 	pin.DisableNotification = false

	// 	bot.PinChatMessage(pin)
	// }
}
