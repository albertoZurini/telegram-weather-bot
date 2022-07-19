package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	weatherhandler "github.com/albertoZurini/telegram-weather-bot/weatherHandlerAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	telegramToken := os.Getenv("TELEGRAM_API_TOKEN")

	// Init Telegram
	bot, err := tgbotapi.NewBotAPI(telegramToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	weather, err := weatherhandler.NewWeatherHandler(os.Getenv("WEATHER_API_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

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
				msg.Text = "This will display the bot's manual."
			default:
				msg.Text = "This is not a valid command."
			}

			if strings.Contains(update.Message.Command(), "getWeather") {
				city := strings.Replace(update.Message.Command(), "getWeather", "", -1)
				wi, err := weather.GetWeatherForLocation(city)

				if err != nil {
					msg.Text = "Error retrieving data."
				} else {
					msg.Text = wi.CurrentWeather
				}
			}

			bot.Send(msg)
		} else {
			text := ""
			wi, err := weather.GetWeatherForLocation(update.Message.Text)

			if err != nil {
				text = err.Error()
			} else {
				text = wi.CurrentWeather
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
