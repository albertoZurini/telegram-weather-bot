package weatherHandlerAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
