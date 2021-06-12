package finalspace3

import (
	"errors"
	gohttp "net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/bgildson/gomock-example/entity"
	"github.com/bgildson/gomock-example/pkg/net/http"
)

func TestClientGetQuotes(t *testing.T) {
	testsCases := []struct {
		name          string
		setupMocks    func(httpClient *http.MockClient)
		clientBaseUrl string
		hasError      bool
		expected      []entity.Quote
	}{
		{
			name: "success",
			setupMocks: func(httpClient *http.MockClient) {
				httpClient.
					EXPECT().
					Do(gomock.Any()).
					Return(
						http.NewResponse(gohttp.StatusOK, `[{"id": 1, "quote": "How about fricken’ no?!", "by": "Gary Goodspeed", "character": "https://finalspaceapi.com/api/v0/character/1", "image": "https://finalspaceapi.com/api/character/avatar/gary_goodspeed.jpg"}]`),
						nil,
					)
			},
			clientBaseUrl: "http://valid.url",
			expected:      []entity.Quote{{ID: 1, Quote: "How about fricken’ no?!", By: "Gary Goodspeed", Character: "https://finalspaceapi.com/api/v0/character/1", Image: "https://finalspaceapi.com/api/character/avatar/gary_goodspeed.jpg"}},
		},
		{
			name: "invalid baseUrl",
			setupMocks: func(httpClient *http.MockClient) {
				httpClient.
					EXPECT().
					Do(gomock.Any()).
					Times(0)
			},
			clientBaseUrl: ":invalid:",
			hasError:      true,
		},
		{
			name: "unexpected error with http client",
			setupMocks: func(httpClient *http.MockClient) {
				httpClient.
					EXPECT().
					Do(gomock.Any()).
					Return(nil, errors.New("unexpected error"))
			},
			clientBaseUrl: "http://valid.url",
			hasError:      true,
		},
		{
			name: "could not reach the server",
			setupMocks: func(httpClient *http.MockClient) {
				httpClient.
					EXPECT().
					Do(gomock.Any()).
					Return(
						http.NewResponse(gohttp.StatusGatewayTimeout, ""),
						nil,
					)
			},
			clientBaseUrl: "http://valid.url",
			hasError:      true,
		},
		{
			name: "invalid response body",
			setupMocks: func(httpClient *http.MockClient) {
				httpClient.
					EXPECT().
					Do(gomock.Any()).
					Return(
						http.NewResponse(gohttp.StatusOK, "invalid"),
						nil,
					)
			},
			clientBaseUrl: "http://valid.url",
			hasError:      true,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			httpClient := http.NewMockClient(ctrl)

			tc.setupMocks(httpClient)

			client := New(httpClient, tc.clientBaseUrl)

			quotes, err := client.GetQuotes()
			if tc.hasError && err == nil {
				t.Error("was expecting an error")
			}
			if !reflect.DeepEqual(tc.expected, quotes) {
				t.Errorf("was expecting %+v, but returns %+v", tc.expected, quotes)
			}
		})
	}
}
