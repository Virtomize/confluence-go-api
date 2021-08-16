package goconfluence

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Template contains blueprint data
type Template struct {
	ID          string `json:"templateId,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"templateType,omitempty"`
	Description string `json:"description"`
	Body        Body   `json:"body"`
	Space       Space  `json:"space"`
}

// TemplateQuery defines the query parameters
type TemplateQuery struct {
	SpaceKey string
	Start    int // page start
	Limit    int // page limit
	Expand   []string
}

// TemplateSearch contains blueprint search results
type TemplateSearch struct {
	Results []Template `json:"results"`
	Start   int        `json:"start,omitempty"`
	Limit   int        `json:"limit,omitempty"`
	Size    int        `json:"size,omitempty"`
}

func (a *API) getBlueprintTemplatesEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/template/blueprint")
}

func (a *API) getContentTemplatesEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/template/page")
}

// GetBlueprintTemplates querys for content blueprints defined by TemplateQuery parameters
func (a *API) GetBlueprintTemplates(query TemplateQuery) (*TemplateSearch, error) {
	ep, err := a.getBlueprintTemplatesEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addTemplateQueryParams(query).Encode()

	req, err := http.NewRequest("GET", ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var search TemplateSearch

	err = json.Unmarshal(res, &search)
	if err != nil {
		return nil, err
	}
	return &search, nil
}

// GetContentTemplates querys for content templates
func (a *API) GetContentTemplates(query TemplateQuery) (*TemplateSearch, error) {
	ep, err := a.getContentTemplatesEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addTemplateQueryParams(query).Encode()

	req, err := http.NewRequest("GET", ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var search TemplateSearch

	err = json.Unmarshal(res, &search)
	if err != nil {
		return nil, err
	}
	return &search, nil
}

func addTemplateQueryParams(query TemplateQuery) *url.Values {
	data := url.Values{}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.SpaceKey != "" {
		data.Set("spaceKey", query.SpaceKey)
	}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	return &data
}
