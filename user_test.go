package goconfluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.Nil(t, err)

	u, err := api.CurrentUser()
	assert.Nil(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.AnonymousUser()
	assert.Nil(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.User("42")
	assert.Nil(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.User(":42")
	assert.Nil(t, err)
	assert.Equal(t, &User{}, u)
}
