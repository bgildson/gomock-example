package http

import (
	"bytes"
	"io/ioutil"
	gohttp "net/http"
)

// NewResponse is a short way to create a simple http response
func NewResponse(statusCode int, body string) *gohttp.Response {
	return &gohttp.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     make(gohttp.Header),
	}
}
