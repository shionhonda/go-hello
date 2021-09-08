package main

import (
	"testing"
	"time"
)

func mockFunc(array []string, sep string) string {
	time.Sleep(100 * time.Millisecond)
	return ""
}

func TestCreateStringArray(t *testing.T) {
	n := 3
	array := createStringArray(n)
	m := len(array)
	if m != n {
		t.Errorf("Array length is expected to be %d, but got %d", n, m)
	}
	for _, s := range array {
		if s != "a" {
			t.Errorf("Element is expected to be 'a', but got %s", s)
		}
	}
}

func TestTimer(t *testing.T) {
	secs := timer(mockFunc, []string{"hoge"})
	if secs < 0.09 || secs > 0.11 {
		t.Errorf("Timer is not accurate enough")
	}
}

func TestMyJoin(t *testing.T) {
	array := []string{"hoge", "fuga"}
	s := myJoin(array, "+")
	expected := "hoge+fuga"
	if s != expected {
		t.Errorf("Joined string is expected to be %s, but got %s", expected, s)
	}
}
