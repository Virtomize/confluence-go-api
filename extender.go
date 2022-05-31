package goconfluence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
func (a *API) getAddCategoryEndpoint(spaceKey string, category string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/extender/1.0/category/addSpaceCategory/space/" + spaceKey + "/category/" + category)
}

// AddSpaceCategory /rest/extender/1.0/category/addSpaceCategory/space/{SPACE_KEY}/category/{CATEGORY_NAME}
func (a *API) AddSpaceCategory(spaceKey string, category string) (*AddCategoryResponseType, error) {

	ep, err := a.getAddCategoryEndpoint(spaceKey, category)
	if err != nil {
		return nil, err
	}

	return a.SendAddCategoryRequest(ep, "PUT")

}

func (a *API) SendAddCategoryRequest(ep *url.URL, method string) (*AddCategoryResponseType, error) {

	var addresp AddCategoryResponseType

	err3 := a.DoRequest(ep.String(), method, &addresp)
	if err3 != nil {
		return nil, err3
	}
	return &addresp, nil
}

type PaginationOptions struct {
	// StartAt: The starting index of the returned projects. Base index: 0.
	StartAt int `url:"startAt,omitempty"`
	// MaxResults: The maximum number of projects to return per page. Default: 50.
	MaxResults int `url:"maxResults,omitempty"`
	// Expand: Expand specific sections in the returned issues
}

//type PermissionsTypes []string

func (a *API) GetPermissionTypes() (*PermissionsTypes, error) {
	u := a.endPoint.String() + "/rest/extender/1.0/permission/space/permissionTypes"

	var types PermissionsTypes
	err3 := a.DoRequest(u, "GET", &types)
	if err3 != nil {
		return nil, err3
	}
	return &types, nil
}

func (a *API) GetAllUsersWithAnyPermission(spacekey string, options *PaginationOptions) (*GetAllUsersWithAnyPermissionType, error) {
	var u string = a.endPoint.String() + fmt.Sprintf("/rest/extender/1.0/permission/space/%s/allUsersWithAnyPermission", spacekey)
	/*
		if options == nil {
			u = a.endPoint.String() + fmt.Sprintf("/rest/extender/1.0/permission/space/%s/allUsersWithAnyPermission", spacekey)
		} else {
			u = a.endPoint.String() + fmt.Sprintf("/rest/extender/1.0/permission/space/%s/allUsersWithAnyPermission?startAt=%d&maxResults=%d", spacekey, options.StartAt, options.MaxResults)
		}
	*/
	endpoint, err := addOptions(u, options)
	if err != nil {
		return nil, err
	}
	var types GetAllUsersWithAnyPermissionType

	err = a.DoRequest(endpoint, "GET", &types)
	if err != nil {
		return nil, err
	}
	return &types, nil
}

func (a *API) GetUserPermissionsForSpace(spacekey, user string) (*GetPermissionsForSpaceType, error) {
	u := a.endPoint.String() + fmt.Sprintf("/rest/extender/1.0/permission/user/%s/getPermissionsForSpace/space/%s", user, spacekey)
	var permissions GetPermissionsForSpaceType

	err3 := a.DoRequest(u, "GET", &permissions)
	if err3 != nil {
		return nil, err3
	}
	return &permissions, nil
}

func (a *API) DoRequest(endpoint string, method string, responseContainer interface{}) error {
	if a.Debug {
		fmt.Printf("Send: %s, Method: %s \n", endpoint, method)
	}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}
	resp, err2 := a.Request(req)
	if err2 != nil {
		return err2
	}

	err = json.Unmarshal(resp, &responseContainer)
	if err != nil {
		return err
	}
	if a.Debug {
		fmt.Printf("Reply: %s \n", resp)
	}
	return nil
}

type GetGroupMembersOptions struct {
	MaxResults int `json:"maxResults"`
	StartAt    int `json:"startAt"`
}

type GroupsType struct {
	Total      int      `json:"total"`
	MaxResults int      `json:"maxResults"`
	Groups     []string `json:"groups"`
	StartAt    int      `json:"startAt"`
	Status     string   `json:"status"`
}

//groups, _ := confluence.GetGroups(&gropt)

func (a *API) GetGroups(options *GetGroupMembersOptions) (*GroupsType, error) {

	u := a.endPoint.String() + "/rest/extender/1.0/group/getGroups"

	endpoint, err := addOptions(u, options)
	if err != nil {
		return nil, err
	}

	var groups GroupsType

	err3 := a.DoRequest(endpoint, "GET", &groups)
	if err3 != nil {
		return nil, err3
	}
	return &groups, nil
}

type GetAllGroupsWithAnyPermissionType struct {
	Total      int      `json:"total"`
	MaxResults int      `json:"maxResults"`
	Groups     []string `json:"groups"`
	StartAt    int      `json:"startAt"`
}

func (a *API) GetAllGroupsWithAnyPermission(spacekey string, options *PaginationOptions) (*GetAllGroupsWithAnyPermissionType, error) {

	u := fmt.Sprintf("/rest/extender/1.0/permission/space/%s/allGroupsWithAnyPermission", spacekey)
	endpoint, err := addOptions(u, options)
	if err != nil {
		return nil, err
	}

	groups := new(GetAllGroupsWithAnyPermissionType)
	err3 := a.DoRequest(endpoint, "GET", &groups)
	if err3 != nil {
		return nil, err3
	}

	return groups, nil
}

func (a *API) GetGroupPermissionsForSpace(spacekey, group string) (*GetPermissionsForSpaceType, error) {
	u := fmt.Sprintf("/rest/extender/1.0/permission/group/%s/getPermissionsForSpace/space/%s", group, spacekey)
	permissions := new(GetPermissionsForSpaceType)
	err3 := a.DoRequest(u, "GET", &permissions)
	if err3 != nil {
		return nil, err3
	}

	//defer CleanupH(resp)
	return permissions, nil
}
