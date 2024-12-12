package main

import (
	"github.com/felipedsf/go-samples/deploy_cloudrun/handlers"
	"github.com/felipedsf/go-samples/deploy_cloudrun/internal/integrations"
	log "github.com/sirupsen/logrus"
	"os"

	"net/http"
)

func main() {
	weatherKey := os.Getenv("WEATHER_API_KEY")
	port := os.Getenv("SERVICE_PORT")

	viaCepApi := integrations.NewViaCepApi()
	weatherApi := integrations.NewWeatherApi(weatherKey)

	handler := handlers.NewTemperatureHandler(viaCepApi, weatherApi)
	http.HandleFunc("/{cep}", handler.Check)

	log.Info("Running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
