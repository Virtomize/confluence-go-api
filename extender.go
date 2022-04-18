package goconfluence

import (
	"net/http"
	"net/url"
)

func (a *API) getAddCategoryEndpoint(spaceKey string, category string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/extender/1.0/category/addSpaceCategory/space/" + spaceKey + "/category/" + category)
}

// AddSpaceCategory /rest/extender/1.0/category/addSpaceCategory/space/{SPACE_KEY}/category/{CATEGORY_NAME}
func (a *API) AddSpaceCategory(spaceKey string, category string) (*http.Response, error) {

	ep, err := a.getAddCategoryEndpoint(spaceKey, category)
	if err != nil {
		return nil, err
	}

	return nil, a.SendAddCategoryRequest(ep, "PUT")

}

func (a *API) SendAddCategoryRequest(ep *url.URL, method string) error {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return err
	}

	_, err2 := a.Request(req)
	if err2 != nil {
		return err2
	}

	return nil
}

/*
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
*/
