package goconfluence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type apiTestValue struct {
	Input []string
	Error error
}

func TestNewAPI(t *testing.T) {
	testValues := []apiTestValue{
		{[]string{"", "username", "token"}, fmt.Errorf("url, username or token empty")},
		{[]string{"test", "", "token"}, fmt.Errorf("url, username or token empty")},
		{[]string{"test", "username", ""}, fmt.Errorf("url, username or token empty")},
		{[]string{"https://test.test", "username", "token"}, nil},
		{[]string{"test", "username", "token"}, fmt.Errorf("parse \"test\": invalid URI for request")},
	}

	for _, test := range testValues {
		api, err := NewAPI(test.Input[0], test.Input[1], test.Input[2])
		if err != nil {
			assert.Equal(t, test.Error.Error(), err.Error())
		} else {
			assert.Equal(t, test.Input[0], api.endPoint.String())
			assert.Equal(t, test.Input[1], api.username)
			assert.Equal(t, test.Input[2], api.token)
		}
	}
}

func TestSetDebug(t *testing.T) {
	assert.False(t, DebugFlag)
	SetDebug(true)
	assert.True(t, DebugFlag)
	SetDebug(false)
	assert.False(t, DebugFlag)
}
