package goconfluence

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentTemplatesEndpoints(t *testing.T) {
	a, err := NewAPI("https://test.test", "", "")
	assert.Nil(t, err)

	ep, err := a.getContentTemplatesEndpoint()
	assert.Nil(t, err)
	uri, err := url.ParseRequestURI("https://test.test/template/page")
	assert.Nil(t, err)
	assert.Equal(t, ep, uri)
}

func TestGetBlueprintTemplatesEndpoints(t *testing.T) {
	a, err := NewAPI("https://test.test", "", "")
	assert.Nil(t, err)

	ep, err := a.getBlueprintTemplatesEndpoint()
	assert.Nil(t, err)
	uri, err := url.ParseRequestURI("https://test.test/template/blueprint")
	assert.Nil(t, err)
	assert.Equal(t, ep, uri)
}

func TestBlueprintTemplateGetter(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	b, err := api.GetBlueprintTemplates(TemplateQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &TemplateSearch{}, b)
}

func TestContentTemplateGetter(t *testing.T) {

	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	c, err := api.GetContentTemplates(TemplateQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &TemplateSearch{}, c)
}
