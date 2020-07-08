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
	assert.Equal(t, "/content/test", url.Path)

	url, err = a.getContentEndpoint()
	assert.Nil(t, err)
	assert.Equal(t, "/content/", url.Path)

	url, err = a.getContentChildEndpoint("test", "child")
	assert.Nil(t, err)
	assert.Equal(t, "/content/test/child/child", url.Path)

	url, err = a.getContentGenericEndpoint("test", "child")
	assert.Nil(t, err)
	assert.Equal(t, "/content/test/child", url.Path)
}

func TestContentGetter(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
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

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	l, err := api.AddLabels("42", &[]Label{})
	assert.Nil(t, err)
	assert.Equal(t, &Labels{}, l)
}

func TestDeleteLabels(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	l, err := api.DeleteLabel("42", "test")
	assert.Nil(t, err)
	assert.Equal(t, &Labels{}, l)
}

func TestContent(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
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
