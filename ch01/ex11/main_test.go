package main

import (
	"regexp"
	"sort"
	"testing"
)

func TestFetch(t *testing.T) {
	urls := []string{"http://gopl.io", "http://gopl.not_found"}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	results := []string{}
	for range urls {
		results = append(results, <-ch)
	}
	sort.Strings(results)
	reg1 := `^[0-9.]{4}s\s[0-9\s]{7}\s.+$`
	if !check_regexp(reg1, results[0]) {
		t.Errorf("Regexp not match to: %s", results[0])
	}
	reg2 := "fetch.+"
	if !check_regexp(reg2, results[1]) {
		t.Errorf("Expected an error, but got: %s", results[1])
	}

}

func check_regexp(reg string, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
