package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var tests = []string{
		"sample.jpg",
		"sample.png",
		"sample.gif",
	}
	fmtStrs := []string{jpegStr, pngStr, gifStr}
	for _, test := range tests {
		for _, fmtStr := range fmtStrs {
			f, _ := os.Open("fixture/" + test)
			defer f.Close()
			reader := bufio.NewReader(f)
			var buf bytes.Buffer
			err := convert(reader, &buf, fmtStr)
			assert.NoError(t, err)
		}
		f, _ := os.Open("fixture/" + test)
		defer f.Close()
		reader := bufio.NewReader(f)
		var buf bytes.Buffer
		err := convert(reader, &buf, "webp")
		assert.Error(t, err)
	}
}
