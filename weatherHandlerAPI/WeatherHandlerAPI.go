package weatherHandlerAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/albertoZurini/telegram-weather-bot/userHandler"
	owm "github.com/briandowns/openweathermap"
)

type WeatherHandlerAPI struct {
	Token       string     `json:"token"`
	Client      HTTPClient `json:"-"`
	apiEndpoint string
}

func NewWeatherHandler(token string) (*WeatherHandlerAPI, error) {
	return NewWeatherHandlerWithClient(token, "http://api.weatherstack.com/", &http.Client{})
}

func NewWeatherHandlerWithClient(token, apiEndpoint string, client HTTPClient) (*WeatherHandlerAPI, error) {
	wa := &WeatherHandlerAPI{
		Token:       token,
		Client:      client,
		apiEndpoint: apiEndpoint,
	}

	return wa, nil
}

func (wh *WeatherHandlerAPI) GetWeatherForLocation(location string) (*WeatherInformation, error) {
	endPoint := fmt.Sprintf("%s/current?access_key=%s&query=%s", wh.apiEndpoint, wh.Token, location)

	resp, err := http.Get(endPoint)

	if err != nil {
		return nil, err
	}

	weatherInfo, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	wi := &WeatherInformation{
		CurrentWeather: string(weatherInfo),
	}

	return wi, nil
}

func (wi *WeatherHandlerAPI) GetDailyWeatherForLocationByName(cityName string) map[string]interface{} {
	return map[string]interface{}{"todo": "none"}
}

/*
func (wi *WeatherHandlerAPI) GetDailyWeatherForLocationByLocation(location userHandler.Location) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric",
		location.Coordinates[1], location.Coordinates[0], os.Getenv("OPENWEATHERMAP_API_TOKEN"))

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	weatherInfoString, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var weatherInformation map[string]interface{}
	err = json.Unmarshal([]byte(weatherInfoString), &weatherInformation)
	if err != nil {
		return nil, err
	}

	if weatherInformation["cod"].(string) != "200" {
		return nil, fmt.Errorf("%s", weatherInformation["Message"].(string))
	}

	return weatherInformation, nil
}
*/

func (wi *WeatherHandlerAPI) Get5DaysWeatherForLocationByLocation(location userHandler.Location) (*owm.ForecastWeatherData, error) {
	w, err := owm.NewForecast("5", "C", "EN", os.Getenv("OPENWEATHERMAP_API_TOKEN")) // fahrenheit (imperial) with Russian output
	if err != nil {
		return nil, err
	}

	w.DailyByCoordinates(
		&owm.Coordinates{
			Longitude: location.Coordinates[0],
			Latitude:  location.Coordinates[1],
		}, 5)

	return w, nil
}
