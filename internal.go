package goconfluence

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	errEmptyHTTPClient = errors.New("empty http client")
)

// NewAPI implements API constructor
func NewAPI(location string, username string, token string) (*API, error) {
	if len(location) == 0 {
		return nil, errors.New("url empty")
	}

	u, err := url.ParseRequestURI(location)

	if err != nil {
		return nil, err
	}

	a := new(API)
	a.endPoint = u
	a.token = token
	a.username = username

	// #nosec G402
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	a.Client = &http.Client{Transport: tr}

	return a, nil
}

// NewAPIWithClient creates a new API instance using an existing HTTP client.
// Useful when using oauth or other authentication methods.
func NewAPIWithClient(location string, client *http.Client) (*API, error) {
	u, err := url.ParseRequestURI(location)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, errEmptyHTTPClient
	}

	a := new(API)
	a.endPoint = u
	a.Client = client

	return a, nil
}

// VerifyTLS to enable disable certificate checks
func (a *API) VerifyTLS(set bool) {
	// #nosec G402
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !set},
	}
	a.Client = &http.Client{Transport: tr}
}

// DebugFlag is the global debugging variable
var DebugFlag = false

// SetDebug enables debug output
func SetDebug(state bool) {
	DebugFlag = state
}

// Debug outputs debug messages
func Debug(msg interface{}) {
	if DebugFlag {
		fmt.Printf("%+v\n", msg)
	}
}
