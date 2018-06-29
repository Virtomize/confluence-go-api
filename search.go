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
	"net/url"
	"strconv"
)

// getContentEndpoint creates the correct api endpoint by given id
func (a *API) getSearchEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/search")
}

// Search querys confluence using CQL
func (a *API) Search(query SearchQuery) (*Search, error) {
	ep, err := a.getSearchEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addSearchQueryParams(query).Encode()
	return a.SendSearchRequest(ep, "GET")
}

// addSearchQueryParams adds the defined query parameters
func addSearchQueryParams(query SearchQuery) *url.Values {

	data := url.Values{}
	if query.CQL != "" {
		data.Set("cql", query.CQL)
	}
	if query.CQLContext != "" {
		data.Set("cqlcontext", query.CQLContext)
	}
	if query.IncludeArchivedSpaces != true {
		data.Set("includeArchivedSpaces", "true")
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	return &data
}
