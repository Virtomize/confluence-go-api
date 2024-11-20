package goconfluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAttachmentEndpoint(t *testing.T) {
	a, err := NewAPI("https://test.test", "username", "token")
	assert.Nil(t, err)

	url, err := a.getAttachmentEndpoint("test")
	assert.Nil(t, err)
	assert.Equal(t, "/attachments/test", url.Path)
}

func TestAttachmentGetter(t *testing.T) {
	server := confluenceRestAPIStub()
	defer server.Close()

	api, err := NewAPI(server.URL+"/wiki/api/v2", "userame", "token")
	assert.Nil(t, err)

	c, err := api.GetAttachmentById("2495990589")
	assert.Nil(t, err)
	assert.Equal(t, &Attachment{}, c)
}
