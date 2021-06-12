package finalspace1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bgildson/gomock-example/entity"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
}

func New(httpClient *http.Client, baseUrl string) *Client {
	return &Client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

func (c Client) GetQuotes() ([]entity.Quote, error) {
	res, err := c.httpClient.Get(c.baseUrl + "/api/v0/quote")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not reach the external api")
	}

	var quote []entity.Quote
	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		return nil, err
	}

	return quote, nil
}
