package goconfluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	s, err := api.GetAllSpaces(AllSpacesQuery{})
	assert.Nil(t, err)
	assert.Equal(t, &AllSpaces{}, s)

}
