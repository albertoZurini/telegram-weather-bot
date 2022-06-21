package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	telegram_token := os.Getenv("TELEGRAM_API_TOKEN")
	chat_id := 0

	// BEGIN INIT
	// Init Telegram
	bot, err := tgbotapi.NewBotAPI(telegram_token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	// Read users
	jsonFile, err := ioutil.ReadFile("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic(err)
	}
	var f map[string]interface{}
	if err := json.Unmarshal(jsonFile, &f); err == nil {
		chat_id = int(f["chatID"].(float64))
	}
	fmt.Print(chat_id)

	weather, err := weather_handler.NewWeatherHandler(os.Getenv("WEATHER_API_TOKEN"))

	fmt.Println(weather)

	/*
		// END INIT

		updateConfig := tgbotapi.NewUpdate(0)

		// Tell Telegram we should wait up to 30 seconds on each request for an
		// update. This way we can get information just as quickly as making many
		// frequent requests without having to send nearly as many.
		updateConfig.Timeout = 30

		// Start polling Telegram for updates.
		updates := bot.GetUpdatesChan(updateConfig)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			fmt.Println(int(update.Message.Chat.ID))
			fmt.Println(msg)

			w, err := getCurrent(update.Message.Text, "C", "EN")
			if err != nil {
				log.Fatalln(err)
			}
			/*msg.ReplyToMessageID = update.Message.MessageID
			city := update.Message.Text
			w.CurrentByName(city)*

			fmt.Println(w)
		}*/
}

func weather_handler(s string) {
	panic("unimplemented")
}
