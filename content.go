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
	"strconv"
	"strings"
)

// getContentEndpoint creates the correct api endpoint by given id
func (a *API) getContentEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content/" + id)
}

// GetContent querys content by id
func (a *API) GetContent(id string, query ContentQuery) (*Content, error) {
	ep, err := a.getContentEndpoint(id)
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addContentQueryParams(query).Encode()
	return a.SendContentRequest(ep, "GET", nil)
}

// CreateContent creates content
func (a *API) CreateContent(c *Content) (*Content, error) {
	ep, err := a.getContentEndpoint(c.ID)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "POST", c)
}

// UpdateContent updates content
func (a *API) UpdateContent(c *Content) (*Content, error) {
	ep, err := a.getContentEndpoint(c.ID)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "PUT", c)
}

// DelContent deletes content by id
func (a *API) DelContent(id string) (*Content, error) {
	ep, err := a.getContentEndpoint(id)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "DELETE", nil)
}

// addContentQueryParams adds the defined query parameters
func addContentQueryParams(query ContentQuery) *url.Values {

	data := url.Values{}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.OrderBy != "" {
		data.Set("orderby", query.OrderBy)
	}
	if query.PostingDay != "" {
		data.Set("postingDay", query.PostingDay)
	}
	if query.SpaceKey != "" {
		data.Set("spaceKey", query.SpaceKey)
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	if query.Status != "" {
		data.Set("status", query.Status)
	}
	if query.Title != "" {
		data.Set("title", query.Title)
	}
	if query.Trigger != "" {
		data.Set("trigger", query.Trigger)
	}
	if query.Type != "" {
		data.Set("type", query.Type)
	}
	return &data
}
