package integrations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

const weather_url = "https://api.weatherapi.com/v1/current.json?q=%s&key=%s"

type WeatherApi struct {
	apiKey string
	client http.Client
}

func NewWeatherApi(apiKey string) *WeatherApi {
	return &WeatherApi{apiKey: apiKey, client: http.Client{}}
}

func (w *WeatherApi) GetWeather(ctx context.Context, query string) (*WeatherResponse, error) {
	log.Info("Getting weather for ", query)
	preparedQuery := fmt.Sprintf(weather_url, url.PathEscape(query), w.apiKey)
	req, _ := http.NewRequestWithContext(ctx, "GET", preparedQuery, nil)
	response, err := w.client.Do(req)
	if err != nil {
		return &WeatherResponse{}, errors.New("failed to get weather")
	}
	defer response.Body.Close()
	wBody, err := io.ReadAll(response.Body)
	var weather WeatherResponse
	err = json.Unmarshal(wBody, &weather)
	if err != nil {
		return &WeatherResponse{}, errors.New("failed to unmarshal weather")
	}
	return &weather, nil
}

type WeatherResponse struct {
	Location struct {
		Name      string `json:"name"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}
