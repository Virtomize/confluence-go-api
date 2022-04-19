package goconfluence

import (
	"testing"
)

func Test_TestExtenderAddCategoryResponseType(t *testing.T) {
	prepareTest(t, TestExtenderAddCategoryResponseType)

	ok, err2 := testClient.AddSpaceCategory("ds", "test")
	//	defer CleanupH(resp)
	if err2 == nil {
		if ok == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if ok.Status != "category 'test' added to 'Demonstration Space (ds)' space" {
				t.Errorf("Expected Success, received: %v Spaces \n", ok.Status)
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_ExtenderSpacePermissionTypes(t *testing.T) {
	prepareTest(t, TestExtenderSpacePermissionTypes)

	permissionTypes, err2 := testClient.GetPermissionTypes()
	//	defer CleanupH(resp)
	if err2 == nil {
		if permissionTypes == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(*permissionTypes) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(*permissionTypes))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_TestExtenderSpacePermissionTypes(t *testing.T) {
	prepareTest(t, TestExtenderSpaceUserPermission)

	usersWithAnyPermission, err2 := testClient.GetAllUsersWithAnyPermission("~admin", &PaginationOptions{}) // StartAt: 0, MaxResults: 50
	//	defer CleanupH(resp)
	if err2 == nil {
		if usersWithAnyPermission == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(usersWithAnyPermission.Users) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(usersWithAnyPermission.Users))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_TestExtenderSpaceAnyUserPermission(t *testing.T) {
	prepareTest(t, TestExtenderSpaceAnyUserPermission)

	userPermissionsForSpace, err2 := testClient.GetUserPermissionsForSpace("~admin", "admin")
	//	defer CleanupH(resp)
	if err2 == nil {
		if userPermissionsForSpace == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(userPermissionsForSpace.Permissions) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(userPermissionsForSpace.Permissions))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}
