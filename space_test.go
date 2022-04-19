package goconfluence

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the Jira client being tested.
	testClient *API

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

func TestAllSpacesQueryParams(t *testing.T) {
	query := AllSpacesQuery{
		Expand:           []string{"a", "b"},
		Favourite:        true,
		FavouriteUserKey: "key1",
		SpaceKey:         "KEY",
		Status:           "sta",
		Type:             "global",
		Limit:            1,
		Start:            1,
	}
	p := addAllSpacesQueryParams(query)
	assert.Equal(t, p.Get("expand"), "a,b")
	assert.Equal(t, p.Get("favourite"), "true")
	assert.Equal(t, p.Get("favouriteUserKey"), "key1")
	assert.Equal(t, p.Get("spaceKey"), "KEY")
	assert.Equal(t, p.Get("status"), "sta")
	assert.Equal(t, p.Get("type"), "global")
	assert.Equal(t, p.Get("limit"), "1")
	assert.Equal(t, p.Get("start"), "1")
}

func TestGetAllSpacesQuery(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki", "userame", "token")
	assert.Nil(t, err)

	s, err := api.GetAllSpaces(AllSpacesQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &AllSpaces{}, s)

}

func Test_SpaceGetSpacesMocFileSuccess(t *testing.T) {
	index := TestSpaceGetSpacesMocFileS

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

	spaces, err2 := testClient.GetAllSpaces(AllSpacesQuery{})
	//	defer CleanupH(resp)

	if err2 == nil {

		if spaces == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if spaces.Size != 1 {
				t.Errorf("Expected 1 Space, received: %v Spaces \n", spaces.Size)
			}
		}
	} else {
		t.Error("Received nil response.")
	}

}

// setup sets up a test HTTP server along with a jira.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	var err error

	testClient, err = NewAPI(testServer.URL, "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	testClient.Debug = true

}

// teardown closes the test HTTP server.
func teardown() {
	//	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}
