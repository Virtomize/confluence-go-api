package goconfluence

import (
	"testing"

	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/perolo/gojson"
	"github.com/stretchr/testify/assert"
)

type ConfluenceTestType struct {
	APIEndpoint  string
	MockEndpoint string
	File         string
	Method       string
	Type         string
	TypeFile     string
}

const (
	TestSpaceGetSpacesMocFileS = iota
	TestSpaceGetPersonalSpaces
	TestGetVersion
	TestExtenderAddCategoryResponseType
	TestExtenderSpacePermissionTypes
	TestExtenderSpaceUserPermission
	TestExtenderSpaceAnyUserPermission
	TestExtenderGetGroups
	TestExtenderGetUsers
)

var ConfluenceTest = []ConfluenceTestType{
	{
		MockEndpoint: "/rest/api/space",
		APIEndpoint:  "/rest/api/space", File: "mocks/spaces.json",
		Method: "GET", Type: "AllSpaces", TypeFile: "space-dtos.go",
	},
	{
		MockEndpoint: "/rest/api/space",
		APIEndpoint:  "/rest/api/space?limit=50&status=current&type=personal",
		File:         "mocks/get-permissions.json", Method: "GET", Type: "", TypeFile: "",
	},
	{
		MockEndpoint: "/rest/experimental/content/98319/version",
		APIEndpoint:  "/rest/experimental/content/98319/version", File: "mocks/version.json",
		Method: "GET", Type: "ContentVersionResult", TypeFile: "version-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/category/addSpaceCategory/space/ds/category/test",
		APIEndpoint:  "/rest/extender/1.0/category/addSpaceCategory/space/ds/category/test",
		File:         "mocks/extender-add.json", Method: "PUT", Type: "AddCategoryResponseType",
		TypeFile: "extender-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/permission/space/permissionTypes",
		APIEndpoint:  "/rest/extender/1.0/permission/space/permissionTypes",
		File:         "mocks/permissions.json", Method: "GET", Type: "PermissionsTypes",
		TypeFile: "permissions-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/permission/space/~admin/allUsersWithAnyPermission",
		APIEndpoint:  "/rest/extender/1.0/permission/space/~admin/allUsersWithAnyPermission?maxResults=50",
		File:         "mocks/get-users-permissions.json", Method: "GET", Type: "GetAllUsersWithAnyPermissionType",
		TypeFile: "get-users-permissions-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/permission/user/admin/getPermissionsForSpace/space/~admin",
		APIEndpoint:  "/rest/extender/1.0/permission/user/admin/getPermissionsForSpace/space/~admin",
		File:         "mocks/get-admin-permissions.json", Method: "GET", Type: "GetPermissionsForSpaceType",
		TypeFile: "get-admin-permissions-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/group/getGroups",
		APIEndpoint:  "/rest/extender/1.0/group/getGroups",
		File:         "mocks/get-groups.json",
		Method:       "GET",
		Type:         "GroupsType",
		TypeFile:     "get-groups-dtos.go",
	},
	{
		MockEndpoint: "/rest/extender/1.0/group/getUsers/confluence-users",
		APIEndpoint:  "/rest/extender/1.0/group/getUsers/confluence-users?showExtendedDetails=true",
		File:         "mocks/get-users.json",
		Method:       "GET",
		Type:         "UsersType",
		TypeFile:     "get-users-dtos.go",
	},
}

func UpdateTests() error {
	//	confClient, err := NewAPI("http://localhost:1990/confluence", "admin", "admin")
	confClient, err := NewAPI("http://192.168.50.40:8090", "admin", "admin")
	confClient.Debug = true
	if err != nil {
		return err
	}
	//Remove all old files
	for _, ctest := range ConfluenceTest {
		e := os.Remove(ctest.File)
		if e != nil {
			fmt.Printf("Expected? %s\n", e.Error())
		}
		e = os.Remove(ctest.TypeFile)
		if e != nil {
			fmt.Printf("Expected? %s\n", e.Error())
		}
		resp, err2 := confClient.SendGenericRequest(ctest.APIEndpoint, ctest.Method)
		if err2 != nil {
			return err2
		}
		err3 := ioutil.WriteFile(ctest.File, resp, 0644)
		if err3 != nil {
			return err3
		}
		if ctest.Type != "" {
			i := strings.NewReader(string(resp))
			res, err2 := gojson.Generate(i, gojson.ParseJson, ctest.Type, "goconfluence", []string{"json"}, false, true)
			if err2 != nil {
				return err2
			}
			err3 = ioutil.WriteFile(ctest.TypeFile, res, 0644)
			if err3 != nil {
				return err3
			}
		}
	}
	return nil
}

/*
Requires confluence server up and running...
removed for now...
Add "t_" for now
*/

func TesUpdate(t *testing.T) {

	err := UpdateTests()
	assert.Nil(t, err)

}
