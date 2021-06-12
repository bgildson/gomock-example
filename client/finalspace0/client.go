package finalspace0

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bgildson/gomock-example/entity"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c Client) GetQuotes() ([]entity.Quote, error) {
	res, err := http.DefaultClient.Get("https://finalspaceapi.com/api/v0/quote")
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
