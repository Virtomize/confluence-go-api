package goconfluence

import (
	"net/http"
	"net/url"
)

// API is the main api data structure
type API struct {
	EndPoint        *url.URL
	Client          *http.Client
	Username, Token string
}
