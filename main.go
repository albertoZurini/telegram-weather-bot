package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/albertoZurini/telegram-weather-bot/MessageBeautifier"
	"github.com/albertoZurini/telegram-weather-bot/utils"

	userHandler "github.com/albertoZurini/telegram-weather-bot/userHandler"
	weatherhandler "github.com/albertoZurini/telegram-weather-bot/weatherHandlerAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	utils.SetupLogger()
	logger := utils.Logger

	telegramToken := os.Getenv("TELEGRAM_API_TOKEN")

	// Init Telegram
	bot, err := tgbotapi.NewBotAPI(telegramToken)

	if err != nil {
		logger.Panic(err.Error())
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	weather, err := weatherhandler.NewWeatherHandler(os.Getenv("WEATHER_API_TOKEN"))

	if err != nil {
		logger.Panic(err.Error())
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

		if update.Message.Location != nil {
			latLng := fmt.Sprintf("%f,%f", update.Message.Location.Latitude, update.Message.Location.Longitude)
			wi, err := weather.GetWeatherForLocation(latLng)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if err == nil {
				msg.Text = wi.CurrentWeather
			} else {
				msg.Text = "Error"
			}

			bot.Send(msg)
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyToMessageID = update.Message.MessageID

			switch update.Message.Command() {
			case "id":
				msg.Text = fmt.Sprintf("Chat ID: %d", update.Message.Chat.ID)
			case "help":
				msg.Text = "This will display the bot's manual."
			case "getWeather":
				dbh, err := userHandler.NewDBHandler()
				if err != nil {
					fmt.Print(err)
				}

				loc, err := dbh.GetLocationForUser(update.Message.Chat.ID)
				wi, err := weather.Get5DaysWeatherForLocationByLocation(loc)

				asd := MessageBeautifier.BeautifyDailyWeatherMessage(wi)
				msg.Text = asd

			default:
				if strings.Contains(update.Message.Command(), "getWeather") {
					city := strings.Replace(update.Message.Command(), "getWeather", "", -1)
					wi, err := weather.GetWeatherForLocation(city)

					if err != nil {
						msg.Text = "Error retrieving data."
					} else {
						msg.Text = wi.CurrentWeather
					}
				} else {
					msg.Text = "This is not a valid command."
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
