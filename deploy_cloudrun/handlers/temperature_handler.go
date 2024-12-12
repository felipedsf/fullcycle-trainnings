package handlers

import (
	"context"
	"encoding/json"
	"github.com/felipedsf/go-samples/deploy_cloudrun/internal/integrations"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type TemperatureHandler struct {
	viaCepApi  *integrations.ViaCepApi
	weatherApi *integrations.WeatherApi
}

func NewTemperatureHandler(viaCepApi *integrations.ViaCepApi, weatherApi *integrations.WeatherApi) *TemperatureHandler {
	return &TemperatureHandler{viaCepApi: viaCepApi, weatherApi: weatherApi}
}

func (h *TemperatureHandler) Check(writer http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	viaCepResponse, err := h.viaCepApi.GetCity(ctx, cep)
	if err != nil {
		log.Errorf("error getting city: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	weather, err := h.weatherApi.GetWeather(ctx, viaCepResponse.Localidade)
	if err != nil {
		log.Errorf("error getting weather: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	responseData, err := json.Marshal(
		map[string]string{
			"temp_C": strconv.FormatFloat(weather.Current.TempC, 'f', 1, 64),
			"temp_F": strconv.FormatFloat(weather.Current.TempF, 'f', 1, 64),
			"temp_K": strconv.FormatFloat(weather.Current.TempC+273, 'f', 1, 64),
		})
	if err != nil {
		log.Errorf("error marshaling response: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseData)
}
