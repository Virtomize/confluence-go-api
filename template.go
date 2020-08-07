package goconfluence

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (a *API) getBlueprintTemplatesEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/template/blueprint")
}

func (a *API) getContentTemplatesEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/template/page")
}

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
