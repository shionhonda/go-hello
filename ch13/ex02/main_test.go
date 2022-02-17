package main

import (
	"testing"
)

type link struct {
	value string
	tail  *link
}

func TestNearlyEqual(t *testing.T) {
	// cycle
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	if !IsCycle(a) {
		t.Errorf("%s should have a cycle", a.value)
	}
	if !IsCycle(b) {
		t.Errorf("%s should have a cycle", b.value)
	}
	if !IsCycle(c) {
		t.Errorf("%s should have a cycle", c.value)
	}

	// no cycle
	x, y := &link{value: "x"}, &link{value: "y"}
	x.tail, y.tail = y, nil
	if IsCycle(x) {
		t.Errorf("%s should not have a cycle", x.value)
	}
	if IsCycle(y) {
		t.Errorf("%s should not have a cycle", y.value)
	}
}
