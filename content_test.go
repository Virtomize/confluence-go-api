package goconfluence

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentEndpoints(t *testing.T) {
	a, err := NewAPI("https://test.test", "username", "token")
	assert.Nil(t, err)

	url, err := a.getContentIDEndpoint("test")
	assert.Nil(t, err)
	assert.Equal(t, "/rest/api/content/test", url.Path)

	url, err = a.getContentEndpoint()
	assert.Nil(t, err)
	assert.Equal(t, "/rest/api/content/", url.Path)

	url, err = a.getContentChildEndpoint("test", "child")
	assert.Nil(t, err)
	assert.Equal(t, "/rest/api/content/test/child/child", url.Path)

	url, err = a.getContentGenericEndpoint("test", "child")
	assert.Nil(t, err)
	assert.Equal(t, "/rest/api/content/test/child", url.Path)
}

func TestContentGetter(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki", "userame", "token")
	assert.Nil(t, err)

	c, err := api.GetContentByID("42", ContentQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &Content{}, c)

	s, err := api.GetContent(ContentQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &ContentSearch{}, s)

	p, err := api.GetChildPages("42")
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, p)

	p, err = api.GetComments("42")
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, p)

	p, err = api.GetAttachments("42")
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, p)

	h, err := api.GetHistory("42")
	assert.Nil(t, err)
	assert.Equal(t, &History{}, h)

	l, err := api.GetLabels("42")
	assert.Nil(t, err)
	assert.Equal(t, &Labels{}, l)

	w, err := api.GetWatchers("42")
	assert.Nil(t, err)
	assert.Equal(t, &Watchers{}, w)
}

func TestAddLabels(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki", "userame", "token")
	assert.Nil(t, err)

	l, err := api.AddLabels("42", &[]Label{})
	assert.Nil(t, err)
	assert.Equal(t, &Labels{}, l)
}

func TestDeleteLabels(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki", "userame", "token")
	assert.Nil(t, err)

	l, err := api.DeleteLabel("42", "test")
	assert.Nil(t, err)
	assert.Equal(t, &Labels{}, l)
}

func TestContent(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki", "userame", "token")
	assert.Nil(t, err)

	c, err := api.CreateContent(&Content{})
	assert.Nil(t, err)
	assert.Equal(t, &Content{}, c)

	s, err := api.UploadAttachment("43", "attachmentName", strings.NewReader("attachment content"))
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, s)

	s, err = api.UpdateAttachment("43", "attachmentName", "123", strings.NewReader("attachment content"))
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, s)

	c, err = api.UpdateContent(&Content{})
	assert.Nil(t, err)
	assert.Equal(t, &Content{}, c)

	c, err = api.DelContent("42")
	assert.Nil(t, err)
	assert.Equal(t, &Content{}, c)
}

func TestAddContentQueryParams(t *testing.T) {
	query := ContentQuery{
		Expand:     []string{"foo", "bar"},
		Limit:      1,
		OrderBy:    "test",
		PostingDay: "test",
		SpaceKey:   "test",
		Start:      1,
		Status:     "test",
		Title:      "test",
		Trigger:    "test",
		Type:       "test",
	}

	p := addContentQueryParams(query)

	assert.Equal(t, p.Get("expand"), "foo,bar")
	assert.Equal(t, p.Get("limit"), "1")
	assert.Equal(t, p.Get("orderby"), "test")
	assert.Equal(t, p.Get("postingDay"), "test")
	assert.Equal(t, p.Get("spaceKey"), "test")
	assert.Equal(t, p.Get("start"), "1")
	assert.Equal(t, p.Get("status"), "test")
	assert.Equal(t, p.Get("title"), "test")
	assert.Equal(t, p.Get("trigger"), "test")
	assert.Equal(t, p.Get("type"), "test")
}

func Test_TestGetVersion(t *testing.T) {
	prepareTest(t, TestGetVersion)

	res, err2 := testClient.GetContentVersion("98319")
	//	defer CleanupH(resp)
	if err2 == nil {
		if res == nil {
			t.Error("Expected version - is nil")
		} else {
			if res.Results[0].Number != 1 {
				t.Errorf("Expected Version 1, received: %v \n", res.Results[0].Number)
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

/*
Requires confluence server up and running...
TODO - mock

Add "t_" for now


func TesLocalhost(t *testing.T) {
	//a, err := NewAPI("http://localhost:1990/confluence", "admin", "admin")
	//a, err := NewAPI("http://192.168.50.40:1990/confluence", "admin", "admin")
	a, err := NewAPI("http://192.168.50.40:8090", "admin", "admin")
	assert.Nil(t, err)

	url, err := a.getContentIDEndpoint("test")
	assert.Nil(t, err)
	assert.Equal(t, "/rest/api/content/test", url.Path)

	res, err2 := a.GetPageId("ds", "Welcome to Confluence")
	assert.Nil(t, err2)
	assert.Equal(t, "98319", res.Results[0].ID)

	err3 := a.UppdateAttachment("ds", "Welcome to Confluence", "C:/Temp/Template.xlsx")
	assert.Nil(t, err3)

	a.AddPage("Test Added page 2", "ds", "./mocks/grouppage.html", true, true, res.Results[0].ID)

}
*/
