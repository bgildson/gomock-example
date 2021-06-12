package http

import (
	gohttp "net/http"
)

// Client represents how a http client should be implemented
type Client interface {
	Do(req *gohttp.Request) (*gohttp.Response, error)
}
