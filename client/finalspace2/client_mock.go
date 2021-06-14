package finalspace2

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// RoundTripFunc implements a middleware component that could be
// used to mock requests when using a http Client
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip is the function called to handle the requests
func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

// NewMockedHttpClient creates a new http Client using the RoundTripFunc as middleware
func NewMockedHttpClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// ApiResponse is a shot function the create a basic http Response
func ApiResponse(response string, status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		Header:     make(http.Header),
	}
}
