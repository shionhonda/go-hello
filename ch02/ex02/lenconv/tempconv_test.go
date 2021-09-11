package lenconv

import "testing"

func TestString(t *testing.T) {
	m := Meter(10)
	s := m.String()
	expected := "10m"
	if s != expected {
		t.Errorf("Length String: expected %s, but got %s", expected, s)
	}

	m = Meter(-10.4)
	s = m.String()
	expected = "-10.4m"
	if s != expected {
		t.Errorf("Length String: expected %s, but got %s", expected, s)
	}

	f := Feet(100)
	s = f.String()
	expected = "100ft"
	if s != expected {
		t.Errorf("Length String: expected %s, but got %s", expected, s)
	}

	f = Feet(-20.333)
	s = f.String()
	expected = "-20.333ft"
	if s != expected {
		t.Errorf("Length String: expected %s, but got %s", expected, s)
	}

}
