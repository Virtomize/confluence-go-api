/*
	Go library for atlassian's confluence wiki

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/confluence-go-api
*/

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
