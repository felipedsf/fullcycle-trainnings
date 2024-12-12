package integrations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

const viacep_url = "https://viacep.com.br/ws/%s/json/"

type ViaCepApi struct {
	client http.Client
}

func NewViaCepApi() *ViaCepApi {
	return &ViaCepApi{client: http.Client{}}
}

func (v *ViaCepApi) GetCity(ctx context.Context, query string) (*ViaCepResponse, error) {
	log.Info("Getting city from ViaCep API", " query ", query)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	preparedQuery := fmt.Sprintf(viacep_url, query)
	req, _ := http.NewRequestWithContext(ctx, "GET", preparedQuery, nil)
	res, err := v.client.Do(req)
	if err != nil {
		return &ViaCepResponse{}, errors.New("failed to get ViaCep response")
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	var viaCepResponse ViaCepResponse
	err = json.Unmarshal(b, &viaCepResponse)
	if err != nil {
		return &ViaCepResponse{}, errors.New("failed to unmarshal ViaCep response")
	}
	return &viaCepResponse, nil
}

type ViaCepResponse struct {
	Localidade string `json:"localidade"`
}
