package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"net/http"
	"os"
)

const (
	webHook = "https://blackstitgolangbot.herokuapp.com/"
)

var tgToken = "1394276809:AAEeMvkmK6NUiu2atp0w912SNBAeoKAxv5E"

func main() {
	port := os.Getenv("PORT")

	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal("creation bot: ", err)
	}

	log.Println("bot created")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("Setting webHook %v; error: %v", webHook, err)
	}

	log.Println("webHook set")

	updates := bot.ListenForWebhook("/")
	command := "/vk"
	groupID := "-15365973"

	for update := range updates {
		text := fmt.Sprintf("Command is not %v", command)
		if command != update.Message.Text {
			if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text)); err != nil {
				log.Print(err)
			}
			continue
		}

		items, err := getPosts(groupID)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, item := range items {
			if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, item.Text)); err != nil {
				log.Print(err)
			}
		}

	}

}
