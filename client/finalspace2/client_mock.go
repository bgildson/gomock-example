package finalspace2

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func NewMockedHttpClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func ApiResponse(response string, status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		Header:     make(http.Header),
	}
}
