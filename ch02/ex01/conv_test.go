package tempconv

import "testing"

func TestCToF(t *testing.T) {
	c := Celsius(10)
	f := CToF(c)
	expected := Fahrenheit(50)
	if f != expected {
		t.Errorf("Temp F: expected %.6f, but got %.6f", expected, f)
	}

	c = Celsius(-55.5)
	f = CToF(c)
	expected = Fahrenheit(-67.9)
	if f != expected {
		t.Errorf("Temp F: expected %.6f, but got %.6f", expected, f)
	}
}
