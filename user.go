package goconfluence

import (
	"net/url"
	"strings"
)

// getContentEndpoint creates the correct api endpoint by given id
func (a *API) getUserEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/user/" + id)
}

// CurrentUser return current user information
func (a *API) CurrentUser() (*User, error) {
	ep, err := a.getUserEndpoint("current")
	if err != nil {
		return nil, err
	}
	return a.SendUserRequest(ep, "GET")
}

// AnonymousUser return user information for anonymous user
func (a *API) AnonymousUser() (*User, error) {
	ep, err := a.getUserEndpoint("anonymous")
	if err != nil {
		return nil, err
	}
	return a.SendUserRequest(ep, "GET")
}

// User returns user data for defined query
// query can be accountID or username
func (a *API) User(query string) (*User, error) {
	ep, err := a.getUserEndpoint("")
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	if strings.Contains(query, ":") {
		data.Set("accountId", query)
	} else {
		data.Set("username", query)
	}

	ep.RawQuery = data.Encode()
	return a.SendUserRequest(ep, "GET")
}
