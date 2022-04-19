package goconfluence

import (
	"net/http"
	"net/url"
)

// API is the main api data structure
type API struct {
	endPoint        *url.URL
	Client          *http.Client
	username, token string
	Debug           bool
}
