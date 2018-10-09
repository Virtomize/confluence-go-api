package goconfluence

import (
	"errors"
	"net/http"
	"net/url"
)

// NewAPI implements API constructor
func NewAPI(location string, username string, token string) (*API, error) {
	if len(location) == 0 || len(username) == 0 || len(token) == 0 {
		return nil, errors.New("url, username or token empty")
	}

	u, err := url.ParseRequestURI(location)

	if err != nil {
		return nil, err
	}

	a := new(API)
	a.endPoint = u
	a.token = token
	a.username = username
	a.client = &http.Client{}

	return a, nil
}

// Auth implements basic auth
func (a *API) Auth(req *http.Request) {
	req.SetBasicAuth(a.username, a.token)
}
