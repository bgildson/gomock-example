package finalspace3

import (
	"encoding/json"
	"errors"
	gohttp "net/http"

	"github.com/bgildson/gomock-example/entity"
	"github.com/bgildson/gomock-example/pkg/net/http"
)

// Client represents how finalspace client should be implemented
type Client interface {
	GetQuotes() ([]entity.Quote, error)
}

type client struct {
	httpClient http.Client
	baseUrl    string
}

// New creates a new Client instance
func New(httpClient http.Client, baseUrl string) Client {
	return &client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

// GetQuotes reachs the finalspace api to return a quotes list
func (c client) GetQuotes() ([]entity.Quote, error) {
	req, err := gohttp.NewRequest(gohttp.MethodGet, c.baseUrl+"/api/v0/quote", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != gohttp.StatusOK {
		return nil, errors.New("could not reach the external api")
	}

	var quote []entity.Quote
	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		return nil, err
	}

	return quote, nil
}
