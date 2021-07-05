package main

import (
	"encoding/json"
	"github.com/Ragontar/TGBot_phoneBook/pkg/bot"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Для тестов сборки, переделать
	setup.Init("cmd/config.txt")

	b, err := bot.NewBot(setup.GetCfgSet().ConfigMap["TelegramAPI_token"])
	if err != nil {
		log.Fatal("error creating new bot")
	}

	//err = b.SetWebhook("https://724295c21ffe.ngrok.io")
	err = b.SetWebhook("https://telegram-pb-bot.herokuapp.com/1818272836:AAEJ16PJAKZJVsx3pO2u9kS6laBKacxr9A8")
	if err != nil {
		log.Fatal("error setting webhook")
	}

	ch := make(chan bot.Update, 10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var update bot.Update
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &update)
		if err != nil {
			log.Println("error decoding update")
		}

		ch <- update
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for update := range ch {
		_, err2 := b.SendMessage(update.Message.Chat.Id, update.Message.Text)
		if err2 != nil {
			log.Println("Error sending message")
		}
	}

}
