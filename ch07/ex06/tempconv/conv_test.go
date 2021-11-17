package tempconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCToF(t *testing.T) {
	var tests = []struct {
		c    Celsius
		want Fahrenheit
	}{
		{Celsius(10), Fahrenheit(50)},
		{Celsius(-10), Fahrenheit(14)},
		{Celsius(0), Fahrenheit(32)},
		{Celsius(1.4), Fahrenheit(34.52)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(CToF(test.c)), 0.001)
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		f    Fahrenheit
		want Celsius
	}{
		{Fahrenheit(10), Celsius(-12.2222)},
		{Fahrenheit(-10), Celsius(-23.3333)},
		{Fahrenheit(0), Celsius(-17.7778)},
		{Fahrenheit(1.4), Celsius(-17)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(FToC(test.f)), 0.001)
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		c    Celsius
		want Kelvin
	}{
		{Celsius(10), Kelvin(283.15)},
		{Celsius(-10), Kelvin(263.15)},
		{Celsius(0), Kelvin(273.15)},
		{Celsius(1.4), Kelvin(274.55)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(CToK(test.c)), 0.001)
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		k    Kelvin
		want Celsius
	}{
		{Kelvin(10), Celsius(-263.15)},
		{Kelvin(0), Celsius(-273.15)},
		{Kelvin(1.4), Celsius(-271.75)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(KToC(test.k)), 0.001)
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		f    Fahrenheit
		want Kelvin
	}{
		{Fahrenheit(10), Kelvin(260.928)},
		{Fahrenheit(-10), Kelvin(249.817)},
		{Fahrenheit(32), Kelvin(273.15)},
		{Fahrenheit(1.4), Kelvin(256.15)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(FToK(test.f)), 0.001)
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		k    Kelvin
		want Fahrenheit
	}{
		{Kelvin(10), Fahrenheit(-441.67)},
		{Kelvin(0), Fahrenheit(-459.67)},
		{Kelvin(1.4), Fahrenheit(-457.15)},
	}
	for _, test := range tests {
		assert.InDelta(t, float64(test.want), float64(KToF(test.k)), 0.001)
	}
}
