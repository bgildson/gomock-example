package finalspace2

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/bgildson/gomock-example/entity"
)

func TestClientGetQuotes(t *testing.T) {
	testsCases := []struct {
		name             string
		responseBody     string
		responseStatus   int
		responseHasError bool
		responseError    error
		expected         []entity.Quote
	}{
		{
			name:           "success",
			responseBody:   `[{"id": 1, "quote": "How about fricken’ no?!", "by": "Gary Goodspeed", "character": "https://finalspaceapi.com/api/v0/character/1", "image": "https://finalspaceapi.com/api/character/avatar/gary_goodspeed.jpg"}]`,
			responseStatus: http.StatusOK,
			expected:       []entity.Quote{{ID: 1, Quote: "How about fricken’ no?!", By: "Gary Goodspeed", Character: "https://finalspaceapi.com/api/v0/character/1", Image: "https://finalspaceapi.com/api/character/avatar/gary_goodspeed.jpg"}},
		},
		{
			name:             "unexpected error with http client",
			responseBody:     ``,
			responseStatus:   http.StatusInternalServerError,
			responseHasError: true,
			responseError:    errors.New("unexpected error"),
		},
		{
			name:             "could not reach the server",
			responseBody:     ``,
			responseStatus:   http.StatusBadGateway,
			responseHasError: true,
		},
		{
			name:             "invalid response body",
			responseBody:     `invalid`,
			responseStatus:   http.StatusOK,
			responseHasError: true,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			httpClient := NewMockedHttpClient(func(_ *http.Request) (*http.Response, error) {
				if tc.responseError != nil {
					return nil, tc.responseError
				}
				return ApiResponse(tc.responseBody, tc.responseStatus), nil
			})

			client := New(httpClient, "")

			quotes, err := client.GetQuotes()
			if (tc.responseHasError || tc.responseError != nil) && err == nil {
				t.Errorf("was expecting the error: %v", tc.responseError)
			}
			if !reflect.DeepEqual(tc.expected, quotes) {
				t.Errorf("was expecting %+v, but returns %+v", tc.expected, quotes)
			}
		})
	}
}
