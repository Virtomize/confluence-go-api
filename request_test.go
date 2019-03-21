package goconfluence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testValuesRequest struct {
	Endpoint string
	Body     string
	Error    error
}

func TestRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	testValues := []testValuesRequest{
		{"/test", "\"test\"", nil},
		{"/nocontent", "", nil},
		{"/noauth", "", fmt.Errorf("authentication failed")},
		{"/noservice", "", fmt.Errorf("service is not available: 503 Service Unavailable")},
		{"/internalerror", "", fmt.Errorf("internal server error: 500 Internal Server Error")},
		{"/unknown", "", fmt.Errorf("unknown response status: 408 Request Timeout")},
	}

	for _, test := range testValues {

		req, err := http.NewRequest("GET", api.endPoint.String()+test.Endpoint, nil)
		assert.Nil(t, err)

		b, err := api.Request(req)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error())
		}

		assert.Equal(t, string(b), test.Body)
	}
}

type testValuesContentRequest struct {
	Content *Content
	Method  string
	Body    *Content
	Error   error
}

func TestSendContentRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getContentEndpoint()
	assert.Nil(t, err)

	testValues := []testValuesContentRequest{
		{nil, "GET", &Content{}, nil},
		{&Content{}, "GET", &Content{}, nil},
	}

	for _, test := range testValues {
		b, err := api.SendContentRequest(ep, test.Method, test.Content)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

type testValuesUserRequest struct {
	Method string
	Body   *User
	Error  error
}

func TestSendUserRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getUserEndpoint("42")
	assert.Nil(t, err)

	testValues := []testValuesUserRequest{
		{"GET", &User{}, nil},
	}

	for _, test := range testValues {
		b, err := api.SendUserRequest(ep, test.Method)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

type testValuesSearchRequest struct {
	Method string
	Body   *Search
	Error  error
}

func TestSendSearchRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getSearchEndpoint()
	assert.Nil(t, err)

	testValues := []testValuesSearchRequest{
		{"GET", &Search{}, nil},
	}

	for _, test := range testValues {
		b, err := api.SendSearchRequest(ep, test.Method)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

type testValuesHistoryRequest struct {
	Method string
	Body   *History
	Error  error
}

func TestSendHistoryRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getContentGenericEndpoint("42", "history")
	assert.Nil(t, err)

	testValues := []testValuesHistoryRequest{
		{"GET", &History{}, nil},
	}

	for _, test := range testValues {
		b, err := api.SendHistoryRequest(ep, test.Method)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

type testValuesLabelRequest struct {
	Method string
	Body   *Labels
	Error  error
	Labels *[]Label
}

func TestSendLabelRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getContentGenericEndpoint("42", "label")
	assert.Nil(t, err)

	testValues := []testValuesLabelRequest{
		{"GET", &Labels{}, nil, &[]Label{}},
	}

	for _, test := range testValues {
		b, err := api.SendLabelRequest(ep, test.Method, test.Labels)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

type testValuesWatcherRequest struct {
	Method string
	Body   *Watchers
	Error  error
}

func TestSendWatcherRequest(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	ep, err := api.getContentGenericEndpoint("42", "notification/child-created")
	assert.Nil(t, err)

	testValues := []testValuesWatcherRequest{
		{"GET", &Watchers{}, nil},
	}

	for _, test := range testValues {
		b, err := api.SendWatcherRequest(ep, test.Method)
		if test.Error == nil {
			assert.Nil(t, err)
		} else {
			assert.Equal(t, test.Error.Error(), err.Error)
		}
		assert.Equal(t, test.Body, b)
	}
}

func confluenceRestAPIStub() *httptest.Server {
	var resp interface{}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/wiki/rest/api/test":
			resp = "test"
		case "/wiki/rest/api/noauth":
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		case "/wiki/rest/api/nocontent":
			http.Error(w, "", http.StatusNoContent)
			return
		case "/wiki/rest/api/noservice":
			http.Error(w, "", http.StatusServiceUnavailable)
			return
		case "/wiki/rest/api/internalerror":
			http.Error(w, "", http.StatusInternalServerError)
			return
		case "/wiki/rest/api/unknown":
			http.Error(w, "", http.StatusRequestTimeout)
			return
		case "/wiki/rest/api/content/":
			resp = Content{}
		case "/wiki/rest/api/content/42":
			resp = Content{}
		case "/wiki/rest/api/user/42":
			resp = User{}
		case "/wiki/rest/api/search":
			resp = Search{}
		case "/wiki/rest/api/content/42/history":
			resp = History{}
		case "/wiki/rest/api/content/42/label":
			resp = Labels{}
		case "/wiki/rest/api/content/42/label/test":
			resp = Labels{}
		case "/wiki/rest/api/content/42/notification/child-created":
			resp = Watchers{}
		case "/wiki/rest/api/content/42/child/page":
			resp = Search{}
		case "/wiki/rest/api/content/42/child/attachment":
			resp = Search{}
		case "/wiki/rest/api/content/42/child/comment":
			resp = Search{}
		case "/wiki/rest/api/content/42/child/history":
			resp = Search{}
		case "/wiki/rest/api/content/42/child/label":
			resp = Search{}
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}))
}
