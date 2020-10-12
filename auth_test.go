package goconfluence

import (
	"encoding/base64"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {

	req := httptest.NewRequest("POST", "https://test.test", nil)

	api, err := NewAPI("https://test.test", "username", "token")

	assert.Nil(t, err)
	assert.Empty(t, req.Header)

	api.Auth(req)
	h := req.Header.Get("Authorization")
	assert.NotEmpty(t, h)

	split := strings.Split(h, " ")
	assert.Len(t, split, 2)

	b, err := base64.StdEncoding.DecodeString(split[1])
	assert.Nil(t, err)

	auth := strings.Split(string(b), ":")
	assert.Len(t, auth, 2)
	assert.Equal(t, "username", auth[0])
	assert.Equal(t, "token", auth[1])
}

func TestEmptyAuth(t *testing.T) {

	req := httptest.NewRequest("POST", "https://test.test", nil)

	api, err := NewAPI("https://test.test", "", "")

	assert.Nil(t, err)
	assert.Empty(t, req.Header)

	api.Auth(req)
	h := req.Header.Get("Authorization")
	assert.Empty(t, h)
}
