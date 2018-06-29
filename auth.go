/*
	Go library for atlassian's confluence wiki

	Copyright (C) 2018 Carsten Seeger

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
	@copyright Copyright (C) 2018 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/confluence-go-api
*/

package goconfluence

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
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

	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	u.Path += "wiki/rest/api"

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
