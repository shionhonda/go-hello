package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostname(t *testing.T) {
	var tests = []struct {
		urlStr   string
		hostname string
	}{
		{"https://www.gopl.io/hoge", "www.gopl.io"},
		{"https://gopl.io/hoge", "gopl.io"},
	}
	for _, test := range tests {
		hostname, err := getHostname(test.urlStr)
		assert.Nil(t, err)
		assert.Equal(t, test.hostname, hostname)
	}
}
