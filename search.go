package goconfluence

import (
	"net/url"
	"strconv"
	"strings"
)

// Search results
type Search struct {
	Results   []Results   `json:"results"`
	Start     int         `json:"start,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	Size      int         `json:"size,omitempty"`
	TotalSize int         `json:"totalSize,omitempty"`
	Links     SearchLinks `json:"_links,omitempty"`
}

// Parsing out the _links section to allow paging etc.
type SearchLinks struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"content,omitempty"`
	Next    string `json:"next,omitempty"`
	Self    string `json:"self,omitempty"`
}

// SearchQuery defines query parameters used for searchng
// Query parameter values https://developer.atlassian.com/cloud/confluence/rest/#api-search-get
type SearchQuery struct {
	CQL                   string
	CQLContext            string
	IncludeArchivedSpaces bool
	Limit                 int
	Start                 int
	Expand                []string
}

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

// Search querys confluence using CQL, with the ability to pass in the Next header
// Empty Next header will run search like normal
func (a *API) SearchWithNext(query SearchQuery, next string) (*Search, error) {
	if next == "" {
		return a.Search(query)
	}
	ep, err := url.ParseRequestURI(a.endPoint.String() + strings.TrimPrefix(next, "/rest/api"))
	if err != nil {
		return nil, err
	}
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
	if query.IncludeArchivedSpaces {
		data.Set("includeArchivedSpaces", "true")
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	return &data
}
