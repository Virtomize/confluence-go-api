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
	"net/url"
	"strings"
)

// getEndpoint creates the correct api endpoint by given id
func (a *API) getEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content"/+id)
}

// GetContent querys content by id
func (a *API) GetContent(id string, query Query) (*Content, error) {
	ep, err := a.getEndpoint(id)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	if query.Type != "" {
		data.Set("type", query.Type)
	}
	if query.Start != 0 {
		data.Set("start", query.Start)
	}
	if query.Limit != 0 {
		data.Set("limit", query.Limit)
	}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	ep.RawQuery = data.Encode()

	return a.SendContentRequest(ep, "GET", nil)
}

// UpdateContent updates content
func (a *API) UpdateContent(c *Content) (*Content, error) {
	ep, err := a.getEndpoint(c.ID)
	if err != nil {
		return nil, err
	}

	return SendContentRequest(ep, "PUT", c)
}

// DelContent deletes content by id
func (a *API) DelContent(id string) (*Content, error) {
	ep, err := a.getEndpoint(id)
	if err != nil {
		return nil, err
	}

	return a.SendContentRequest(ep, "DELETE", nil)
}
