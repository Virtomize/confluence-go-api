package goconfluence

import (
	"net/url"
	"strconv"
	"strings"
)

// getSpaceEndpoint creates the correct api endpoint
func (a *API) getSpaceEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/space")
}

// GetAllSpaces queries content using a query parameters
func (a *API) GetAllSpaces(query AllSpacesQuery) (*AllSpaces, error) {
	ep, err := a.getSpaceEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addAllSpacesQueryParams(query).Encode()
	return a.SendAllSpacesRequest(ep, "GET")
}

// addAllSpacesQueryParams adds the defined query parameters
func addAllSpacesQueryParams(query AllSpacesQuery) *url.Values {

	data := url.Values{}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	if query.Favourite {
		data.Set("favourite", strconv.FormatBool(query.Favourite))
	}
	if query.FavouriteUserKey != "" {
		data.Set("favouriteUserKey", query.FavouriteUserKey)
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
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
	if query.Type != "" {
		data.Set("type", query.Type)
	}
	return &data
}
