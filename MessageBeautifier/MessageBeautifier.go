package MessageBeautifier

import owm "github.com/briandowns/openweathermap"

func BeautifyDailyWeatherMessage(wi *owm.ForecastWeatherData) string {
	return emojiMap[wi.ForecastWeatherJson.(*owm.Forecast5WeatherData).List[0].Weather[0].Icon].(string)

}
