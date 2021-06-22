package goconfluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/rest/api", "userame", "token")
	assert.NoError(t, err)

	u, err := api.CurrentUser()
	assert.NoError(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.AnonymousUser()
	assert.NoError(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.User("42")
	assert.NoError(t, err)
	assert.Equal(t, &User{}, u)

	u, err = api.User(":42")
	assert.NoError(t, err)
	assert.Equal(t, &User{}, u)
}
