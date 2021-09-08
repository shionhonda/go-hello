package main

import (
	"strings"
	"testing"
)

func TestFormatUrl(t *testing.T) {
	prefix := "http://"
	url_prefixed := "http://gopl.io"
	if !strings.HasPrefix(formatUrl(url_prefixed), prefix) {
		t.Errorf("Perfix %s expected", prefix)
	}

	url_no_prefixed := "gopl.io"
	if !strings.HasPrefix(formatUrl(url_no_prefixed), prefix) {
		t.Errorf("Perfix %s expected", prefix)
	}

}
