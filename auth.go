package goconfluence

import (
	"net/http"
)

// Auth implements basic auth
func (a *API) Auth(req *http.Request) {
	req.SetBasicAuth(a.username, a.token)
}
