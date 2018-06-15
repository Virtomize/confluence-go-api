/*
	Go library for attlassians confluence wiki

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
	"encoding/json"
	"fmt"
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

func (a *API) SendContentRequest(ep *url.Url, method string, c *Content) (*Content, err) {

	body := nil
	if c != nil {
		js, err = json.Marshal(c)
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

	var c Content

	err = json.Unmarshal(res, &content)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
