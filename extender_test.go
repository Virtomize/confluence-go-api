package goconfluence

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_TestAddCategoryResponse(t *testing.T) {
	index := TestAddCategoryResponseType

	testAPIEndpoint := ConfluenceTest[index].APIEndpoint

	raw, err := ioutil.ReadFile(ConfluenceTest[index].File)
	if err != nil {
		t.Error(err.Error())
	}

	setup()
	defer teardown()
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, ConfluenceTest[index].Method)
		testRequestURL(t, r, testAPIEndpoint)

		_, err = fmt.Fprint(w, string(raw))
		if err != nil {
			t.Errorf("Error given: %s", err)
		}

	})

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
