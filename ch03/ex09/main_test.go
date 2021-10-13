package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title    string
	params   map[string]string
	expected map[string]float64
}

func TestParseRequest(t *testing.T) {
	testCases := []testCase{
		{"Default", map[string]string{}, map[string]float64{"x": 0., "y": 0., "mag": 1.}},
		{"Not Default 1", map[string]string{"x": "2"}, map[string]float64{"x": 2., "y": 0., "mag": 1.}},
		{"Not Default 2", map[string]string{"x": "-1.", "y": "0.1", "mag": "2.0"}, map[string]float64{"x": -1., "y": 0.1, "mag": 2.}},
		{"Set to Default", map[string]string{"mag": "-2.0"}, map[string]float64{"x": 0., "y": 0., "mag": 1.}},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://localhost:8000/", nil)
			q := req.URL.Query()
			for k, v := range tc.params {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			x, y, mag := parseRequest(req)
			assert.Equal(t, tc.expected["x"], x)
			assert.Equal(t, tc.expected["y"], y)
			assert.Equal(t, tc.expected["mag"], mag)
		})
	}
}
