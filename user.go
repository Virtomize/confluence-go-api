package goconfluence

import (
	"net/url"
	"strings"
)

// User defines user informations
type User struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	UserKey     string `json:"userKey"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

// getUserEndpoint creates the correct api endpoint by given id
func (a *API) getUserEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/user?accountId=" + id)
}

// getCurrentUserEndpoint creates the correct api endpoint by given id
func (a *API) getCurrentUserEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/user/current")
}

// getAnonymousUserEndpoint creates the correct api endpoint by given id
func (a *API) getAnonymousUserEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/user/anonymous")
}

// CurrentUser return current user information
func (a *API) CurrentUser() (*User, error) {
	ep, err := a.getCurrentUserEndpoint()
	if err != nil {
		return nil, err
	}
	return a.SendUserRequest(ep, "GET")
}

// AnonymousUser return user information for anonymous user
func (a *API) AnonymousUser() (*User, error) {
	ep, err := a.getAnonymousUserEndpoint()
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
		lookupParams := strings.Split(query, ":")
		if lookupParams[0] == "key" {
			data.Set("key", lookupParams[1])
		} else {
			data.Set("accountId", lookupParams[1])
		}
	} else {
		data.Set("username", query)
	}

	ep.RawQuery = data.Encode()
	return a.SendUserRequest(ep, "GET")
}
