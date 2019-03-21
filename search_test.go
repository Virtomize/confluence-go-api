package goconfluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchQueryParams(t *testing.T) {
	query := SearchQuery{
		CQL:                   "test",
		CQLContext:            "test",
		IncludeArchivedSpaces: true,
		Limit: 1,
		Start: 1,
	}
	p := addSearchQueryParams(query)
	assert.Equal(t, p.Get("cql"), "test")
	assert.Equal(t, p.Get("cqlcontext"), "test")
	assert.Equal(t, p.Get("includeArchivedSpaces"), "true")
	assert.Equal(t, p.Get("limit"), "1")
	assert.Equal(t, p.Get("start"), "1")
}

func TestSearch(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	s, err := api.Search(SearchQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &Search{}, s)

}
