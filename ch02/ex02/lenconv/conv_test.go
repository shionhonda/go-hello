package lenconv

import "testing"

func TestMToF(t *testing.T) {
	m := Meter(10)
	f := MToF(m)
	expected := Feet(10 * 3.2808)
	if f != expected {
		t.Errorf("Length Feet: expected %.6f, but got %.6f", expected, f)
	}

	m = Meter(-1)
	f = MToF(m)
	expected = Feet(-1 * 3.2808)
	if f != expected {
		t.Errorf("Length Feet: expected %.6f, but got %.6f", expected, f)
	}
}

func TestFToM(t *testing.T) {
	f := Feet(14)
	m := FToM(f)
	expected := Meter(14 / 3.2808)
	if m != expected {
		t.Errorf("Length Meter: expected %.6f, but got %.6f", expected, m)
	}

	f = Feet(-3)
	m = FToM(f)
	expected = Meter(-3 / 3.2808)
	if m != expected {
		t.Errorf("Length Meter: expected %.6f, but got %.6f", expected, m)
	}
}
