package tempconv

import "testing"

func TestString(t *testing.T) {
	c := Celsius(10)
	s := c.String()
	expected := "10˚C"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

	c = Celsius(-10.4)
	s = c.String()
	expected = "-10.4˚C"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

	f := Fahrenheit(100)
	s = f.String()
	expected = "100˚F"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

	f = Fahrenheit(-20.333)
	s = f.String()
	expected = "-20.333˚F"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

	k := Kelvin(300)
	s = k.String()
	expected = "300K"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

	k = Kelvin(0.1)
	s = k.String()
	expected = "0.1K"
	if s != expected {
		t.Errorf("Temp String: expected %s, but got %s", expected, s)
	}

}
