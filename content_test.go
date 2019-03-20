package goconfluence

import (
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

func TestAddContentQueryParams(t *testing.T) {
	// tbd
}
