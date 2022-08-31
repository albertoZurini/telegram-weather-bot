package MessageBeautifier

import (
	"fmt"
	"time"

	owm "github.com/briandowns/openweathermap"
)

type DailyForecast struct {
	currentDay []WeatherInfo `json:"currentDay"`
	nextDays   []WeatherInfo `json:"nextDays"`
}

type WeatherInfo struct {
	timeStamp   int64
	temperature float32
	icon        string
	summary     string
	description string
}

var emojiMap map[string]interface{} = map[string]interface{}{
	"01d": "☀️",
	"02d": "⛅️",
	"03d": "☁️",
	"04d": "☁️",
	"09d": "🌧",
	"10d": "🌦",
	"11d": "⛈",
	"13d": "❄️",
	"50d": "🌫",
	"01n": "🌑",
	"02n": "🌑 ☁",
	"03n": "☁️",
	"04n": "️️☁☁",
	"09n": "🌧",
	"10n": "☔️",
	"11n": "⛈",
	"13n": "❄️",
	"50n": "🌫",
}

func OpenWeatherMapToDailyForecast(wi *owm.ForecastWeatherData) {
	var dailyForecast DailyForecast
	today := time.Now()

	fmt.Print(dailyForecast, today)
	/*
		for _, wiowm := range wi["list"].([]interface{}) {
			tm := time.Now() //time.Unix(wiowm["dt"].(int64), 0)
			fmt.Println(wiowm)

			if tm.Day() == today.Day() {
				dailyForecast.currentDay = append(dailyForecast.currentDay, WeatherInfo{
					timeStamp: tm.Unix(),
				})
			}
		}
	*/
}
