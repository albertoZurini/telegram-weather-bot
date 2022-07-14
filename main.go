package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/albertoZurini/telegram-weather-bot/weather_handler"
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

	//fmt.Println(weather.GetWeatherForLocation("Udine"))

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

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyToMessageID = update.Message.MessageID
			switch update.Message.Command() {
			case "id":
				msg.Text = fmt.Sprintf("Chat ID: %d", update.Message.Chat.ID)
			case "help":
				msg.Text = "Can't help you"
			default:
				msg.Text = "This is not a command"
			}

			if strings.Contains(update.Message.Command(), "getWeather") {
				city := strings.Replace(update.Message.Command(), "getWeather", "", -1)

				wi, err := weather.GetWeatherForLocation(city)

				if err != nil {
					msg.Text = "Error retrieving data"
				} else {
					msg.Text = wi.CurrentWeather
				}
			}

			bot.Send(msg)
		} else {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			wi, err := weather.GetWeatherForLocation(update.Message.Text)

			if err == nil {
				msg.Text = wi.CurrentWeather
			} else {
				msg.Text = err.Error()
			}

			fmt.Println(int(update.Message.Chat.ID))
			fmt.Println(msg)
			/*
				w, err := getCurrent(update.Message.Text, "C", "EN")
				if err != nil {
					log.Fatalln(err)
				}
				/*msg.ReplyToMessageID = update.Message.MessageID
				city := update.Message.Text
				w.CurrentByName(city)*/

			bot.Send(msg)
		}

		//fmt.Println(w)
	}
}
