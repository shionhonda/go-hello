package params

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func TestPackAndUnpack(t *testing.T) {
	testCases := []data{
		{Labels: []string(nil), MaxResults: 0, Exact: false},
		{Labels: []string{"golang", "programming"}, MaxResults: 0, Exact: false},
		{Labels: []string{"golang", "programming"}, MaxResults: 100, Exact: false},
		{Labels: []string{"golang", "programming"}, MaxResults: 0, Exact: true},
	}
	for _, tc := range testCases {
		req, _ := http.NewRequest("GET", "http://localhost:8000/", nil)
		url, _ := Pack(req, &tc)
		var recon data
		req, _ = http.NewRequest("GET", url, nil)
		err := Unpack(req, &recon)
		assert.Nil(t, err)
		assert.EqualValues(t, tc, recon)

	}
}
