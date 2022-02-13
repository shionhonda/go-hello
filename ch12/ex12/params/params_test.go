package params

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max,default=10"`
	Exact      bool     `http:"x"`
}

func TestUnpack(t *testing.T) {
	type testCase struct {
		url      string
		expected data
	}
	testCases := []testCase{
		{"http://localhost:8000/", data{Labels: []string(nil), MaxResults: 10, Exact: false}},
		{"http://localhost:8000/?l=golang&l=programming", data{Labels: []string{"golang", "programming"}, MaxResults: 10, Exact: false}},
		{"http://localhost:8000/?l=golang&l=programming&max=100", data{Labels: []string{"golang", "programming"}, MaxResults: 100, Exact: false}},
		{"http://localhost:8000/?x=true&l=golang&l=programming", data{Labels: []string{"golang", "programming"}, MaxResults: 10, Exact: true}},
	}
	for _, tc := range testCases {
		req, _ := http.NewRequest("GET", tc.url, nil)
		var params data
		err := Unpack(req, &params)
		assert.Nil(t, err)
		assert.EqualValues(t, tc.expected, params)

	}
}
