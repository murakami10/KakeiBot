package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "John."
	fmt.Fprint(w, "Hi there, I'm ", text)
}

func main() {

	//bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("ACCESS_TOKEN"))
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//message := linebot.NewTextMessage("hello, world")
	//
	//if _, err := bot.BroadcastMessage(message).Do(); err != nil {
	//	log.Fatal()
	//}

	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
