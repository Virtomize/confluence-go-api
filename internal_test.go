package goconfluence

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type apiTestValue struct {
	Name  string
	Input []string
	Error error
}

type apiClientTestValue struct {
	Name     string
	Location string
	Client   *http.Client
	Error    error
}

func TestNewAPI(t *testing.T) {
	assert := assert.New(t)

	testValues := []apiTestValue{
		{"empty-url", []string{"", "username", "token"}, fmt.Errorf("url empty")},
		{"no-auth", []string{"https://test.test", "", ""}, nil},
		{"basic-auth", []string{"https://test.test", "username", "token"}, nil},
		{"invalid-url", []string{"test", "username", "token"}, fmt.Errorf("parse \"test\": invalid URI for request")},
	}

	for _, test := range testValues {
		t.Run(test.Name, func(t *testing.T) {
			api, err := NewAPI(test.Input[0], test.Input[1], test.Input[2])
			if err != nil {
				assert.Equal(test.Error.Error(), err.Error())
			} else {
				assert.Equal(test.Input[0], api.endPoint.String())
				assert.Equal(test.Input[1], api.username)
				assert.Equal(test.Input[2], api.token)
			}
		})
	}

	testClientValues := []apiClientTestValue{
		{"valid-client", "https://test.test", &http.Client{}, nil},
		{"nil-client", "https://test.test", nil, errEmptyHTTPClient},
		{"invalid-url", "no-url", &http.Client{}, fmt.Errorf("parse \"no-url\": invalid URI for request")},
	}

	for _, test := range testClientValues {
		t.Run(test.Name, func(t *testing.T) {
			api, err := NewAPIWithClient(test.Location, test.Client)
			if err != nil {
				assert.Equal(test.Error.Error(), err.Error())
			} else {
				assert.Equal(api.Client, test.Client)
			}

		})
	}
}

func TestVerifyTLS(t *testing.T) {
	assert := assert.New(t)
	api, err := NewAPI("https://test.test", "", "")
	assert.NoError(err)

	t.Run("set-true", func(t *testing.T) {
		api.VerifyTLS(true)
		assert.Equal(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}, api.Client.Transport)
	})

	t.Run("set-false", func(t *testing.T) {
		api.VerifyTLS(false)
		assert.Equal(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, api.Client.Transport)
	})

}

func TestSetDebug(t *testing.T) {
	assert.False(t, DebugFlag)
	SetDebug(true)
	assert.True(t, DebugFlag)
	SetDebug(false)
	assert.False(t, DebugFlag)
}
