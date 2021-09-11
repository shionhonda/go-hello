package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var x uint64 = 127
	e := 7
	y := PopCount(x)
	if y != e {
		t.Errorf("Expected %d, got %d", e, y)
	}

	x = 1234567890
	e = 12
	y = PopCount(x)
	if y != e {
		t.Errorf("Expected %d, got %d", e, y)
	}
}

func TestPopCountTrick(t *testing.T) {
	var x uint64 = 127
	e := 7
	y := PopCountTrick(x)
	if y != e {
		t.Errorf("Expected %d, got %d", e, y)
	}

	x = 1234567890
	e = 12
	y = PopCountTrick(x)
	if y != e {
		t.Errorf("Expected %d, got %d", e, y)
	}
}
