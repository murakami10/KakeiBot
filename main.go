package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("ACCESS_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage("hello, world")

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal()
	}

}
