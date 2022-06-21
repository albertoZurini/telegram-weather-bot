package weather_handler

import (
	"net/http"
)

type WeatherHandlerAPI struct {
	Token string `json:"token"`

	Client HTTPClient `json:"-"`

	apiEndpoint string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
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

type WeatherInformation struct {
	currentWeather string
}

func (bot *WeatherHandlerAPI) GetWeatherForLocation(location string) (*WeatherInformation, error) {
	wi := &WeatherInformation{
		currentWeather: "diocane",
	}

	return wi, nil
}
