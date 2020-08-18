package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"net/http"
	"os"

)

const (
	webHook = "https://blackstitgolangbot.herokuapp.com/"
)
var tgToken = "1394276809:AAEeMvkmK6NUiu2atp0w912SNBAeoKAxv5E"

func main()  {
	port := os.Getenv("PORT")

	go func(){
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil{
		log.Fatal("creation bot: ", err)
	}

	log.Println("bot created")

	if  _,err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("Setting webHook %v; error: %v", webHook, err,)
	}

	log.Println("webHook set")

	updates := bot.ListenForWebhook("/")
	for update := range updates{
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil{
			log.Print(err)
		}
	}



}