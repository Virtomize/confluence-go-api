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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request implements the basic Request function
func (a *API) Request(req *http.Request) ([]byte, error) {
	req.Header.Add("Accept", "application/json, */*")
	a.Auth(req)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:
		res, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service is not available: %s", resp.Status)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error: %s", resp.Status)
	}

	return nil, fmt.Errorf("unknown response status: %s", resp.Status)
}

// SendContentRequest sends content related requests
// this function is used for getting, updating and deleting content
func (a *API) SendContentRequest(ep *url.URL, method string, c *Content) (*Content, error) {

	var body io.Reader
	if c != nil {
		js, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(js))
	}

	req, err := http.NewRequest(method, ep.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var content Content

	err = json.Unmarshal(res, &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

// SendUserRequest sends user related requests
func (a *API) SendUserRequest(ep *url.URL, method string) (*User, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var user User

	err = json.Unmarshal(res, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SendSearchRequest sends search related requests
func (a *API) SendSearchRequest(ep *url.URL, method string) (*Search, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var search Search

	err = json.Unmarshal(res, &search)
	if err != nil {
		return nil, err
	}

	return &search, nil
}

// SendHistoryRequest requests history
func (a *API) SendHistoryRequest(ep *url.URL, method string) (*History, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var history History

	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}

// SendLabelRequest requests history
func (a *API) SendLabelRequest(ep *url.URL, method string, labels *[]Label) (*Labels, error) {

	var body io.Reader
	if labels != nil {
		js, err := json.Marshal(labels)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(js))
	}

	req, err := http.NewRequest(method, ep.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	if res != nil {
		var l Labels

		err = json.Unmarshal(res, &l)
		if err != nil {
			return nil, err
		}

		return &l, nil
	}

	return &Labels{}, nil
}

// SendWatcherRequest requests watchers
func (a *API) SendWatcherRequest(ep *url.URL, method string) (*Watchers, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}
	var watchers Watchers

	err = json.Unmarshal(res, &watchers)
	if err != nil {
		return nil, err
	}

	return &watchers, nil
}
