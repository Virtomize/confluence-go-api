package goconfluence

import (
	"net/http"
)

// Auth implements basic auth
func (a *API) Auth(req *http.Request) {
	//Supports unauthenticated access to confluence:
	//if username and token are not set, do not add authorization header
	if a.username != "" && a.token != "" {
		req.SetBasicAuth(a.username, a.token)
	} else if a.token != "" {
		req.Header.Set("Authorization", "Bearer "+a.token)
	}
}
