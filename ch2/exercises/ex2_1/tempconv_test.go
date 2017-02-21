// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.1, p. 42: add types, constants, and functions to tempconv for
// processing temperatures in the Kelvin scale, where zero Kelvin is -273.15°C
// and a difference of 1K has the same magnitude as 1°C.
// Based on tempconv.

// Run "go test tempconv_test.go".
package tempconv_test

import (
	"math"
	"testing"

	tc "./tempconv"
)

const Epsilon = 1e-13

func TestCToFToC(t *testing.T) {
	var testdata = []struct {
		c tc.Celsius
		f tc.Fahrenheit
	}{
		{tc.Celsius(tc.AbsoluteZeroC), tc.Fahrenheit(tc.AbsoluteZeroF)},
		{tc.Celsius(tc.FreezingC), tc.Fahrenheit(tc.FreezingF)},
		{tc.Celsius(tc.BoilingC), tc.Fahrenheit(tc.BoilingF)},
	}

	for _, data := range testdata {
		fExp := tc.CToF(data.c)
		if math.Abs(float64(fExp-data.f)) > Epsilon {
			t.Errorf("CToF(%q) = %v; want %v", data.c, fExp, data.f)
		}
		cExp := tc.FToC(data.f)
		if math.Abs(float64(cExp-data.c)) > Epsilon {
			t.Errorf("FToC(%q) = %v; want %v", data.f, cExp, data.c)
		}
	}
}

func TestFToKToF(t *testing.T) {
	var testdata = []struct {
		f tc.Fahrenheit
		k tc.Kelvin
	}{
		{tc.Fahrenheit(tc.AbsoluteZeroF), tc.Kelvin(tc.AbsoluteZeroK)},
		{tc.Fahrenheit(tc.FreezingF), tc.Kelvin(tc.FreezingK)},
		{tc.Fahrenheit(tc.BoilingF), tc.Kelvin(tc.BoilingK)},
	}

	for _, data := range testdata {
		kExp := tc.FToK(data.f)
		if math.Abs(float64(kExp-data.k)) > Epsilon {
			t.Errorf("FToK(%q) = %v; want %v", data.f, kExp, data.k)
		}
		fExp := tc.KToF(data.k)
		if math.Abs(float64(fExp-data.f)) > Epsilon {
			t.Errorf("KToF(%q) = %v; want %v", data.k, fExp, data.f)
		}
	}
}

func TestCToKToC(t *testing.T) {
	var testdata = []struct {
		c tc.Celsius
		k tc.Kelvin
	}{
		{tc.Celsius(tc.AbsoluteZeroC), tc.Kelvin(tc.AbsoluteZeroK)},
		{tc.Celsius(tc.FreezingC), tc.Kelvin(tc.FreezingK)},
		{tc.Celsius(tc.BoilingC), tc.Kelvin(tc.BoilingK)},
	}

	for _, data := range testdata {
		kExp := tc.CToK(data.c)
		if math.Abs(float64(kExp-data.k)) > Epsilon {
			t.Errorf("CToK(%q) = %v; want %v", data.c, kExp, data.k)
		}
		cExp := tc.KToC(data.k)
		if math.Abs(float64(cExp-data.c)) > Epsilon {
			t.Errorf("KToC(%q) = %v; want %v", data.k, cExp, data.c)
		}
	}
}
