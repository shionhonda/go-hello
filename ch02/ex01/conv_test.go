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

func TestCToK(t *testing.T) {
	c := Celsius(100)
	k := CToK(c)
	expected := Kelvin(373.15)
	if k != expected {
		t.Errorf("Temp K: expected %.6f, but got %.6f", expected, k)
	}

	c = Celsius(-15.5)
	k = CToK(c)
	expected = Kelvin(257.65)
	if k != expected {
		t.Errorf("Temp K: expected %.6f, but got %.6f", expected, k)
	}
}

func TestFToC(t *testing.T) {
	f := Fahrenheit(14)
	c := FToC(f)
	expected := Celsius(-10)
	if c != expected {
		t.Errorf("Temp C: expected %.6f, but got %.6f", expected, c)
	}

	f = Fahrenheit(-13)
	c = FToC(f)
	expected = Celsius(-25)
	if c != expected {
		t.Errorf("Temp C: expected %.6f, but got %.6f", expected, c)
	}
}

func TestFToK(t *testing.T) {
	f := Fahrenheit(32)
	k := FToK(f)
	expected := Kelvin(273.15)
	if k != expected {
		t.Errorf("Temp K: expected %.6f, but got %.6f", expected, k)
	}

	f = Fahrenheit(23)
	k = FToK(f)
	expected = Kelvin(268.15)
	if k != expected {
		t.Errorf("Temp K: expected %.6f, but got %.6f", expected, k)
	}
}

func TestKToC(t *testing.T) {
	k := Kelvin(373.15)
	c := KToC(k)
	expected := Celsius(100)
	if c != expected {
		t.Errorf("Temp C: expected %.6f, but got %.6f", expected, c)
	}

	k = Kelvin(257.65)
	c = KToC(k)
	expected = Celsius(-15.5)
	if c != expected {
		t.Errorf("Temp C: expected %.6f, but got %.6f", expected, c)
	}
}

func TestKToF(t *testing.T) {
	k := Kelvin(273.15)
	f := KToF(k)
	expected := Fahrenheit(32)
	if f != expected {
		t.Errorf("Temp F: expected %.6f, but got %.6f", expected, f)
	}

	k = Kelvin(268.15)
	f = KToF(k)
	expected = Fahrenheit(23)
	if f != expected {
		t.Errorf("Temp F: expected %.6f, but got %.6f", expected, f)
	}
}
