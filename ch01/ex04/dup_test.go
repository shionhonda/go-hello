package main

import (
	"testing"
)

func TestCountLines(t *testing.T) {
	counts := make(map[string]int)
	srcs := make(map[string][]string)
	countLines("data/a.txt", counts, srcs)

	if counts["hello"] != 2 {
		t.Errorf("Count of 'hello' is expected to be 2, but got %d", counts["hello"])
	}
	if counts["world"] != 1 {
		t.Errorf("Count of 'world' is expected to be 1, but got %d", counts["world"])
	}
	if counts["gopher"] != 1 {
		t.Errorf("Count of 'gopher' is expected to be 1, but got %d", counts["gopher"])
	}
	if counts["hoge"] != 0 {
		t.Errorf("Count of 'hoge' is expected to be 0, but got %d", counts["hoge"])
	}
}
