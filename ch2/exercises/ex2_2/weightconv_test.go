// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run "go test weightconv_test.go".
package weightconv_test

import (
	"math"
	"testing"

	wc "./weightconv"
)

const Epsilon = 1e-13

func TestKgToLbToKg(t *testing.T) {
	var testdata = []struct {
		kg wc.Kilograms
		lb wc.Pounds
	}{
		{wc.Kilograms(0), wc.Pounds(0)},
		{wc.Kilograms(wc.LbKgRatio), wc.Pounds(1)},
		{wc.Kilograms(1), wc.Pounds(wc.KgLbRatio)},
	}

	for _, data := range testdata {
		lbExp := wc.KgToLb(data.kg)
		if math.Abs(float64(lbExp-data.lb)) > Epsilon {
			t.Errorf("KgToLb(%q) = %v; want %v", data.kg, lbExp, data.lb)
		}
		kgExp := wc.LbToKg(data.lb)
		if math.Abs(float64(kgExp-data.kg)) > Epsilon {
			t.Errorf("LbToKg(%q) = %v; want %v", data.lb, kgExp, data.kg)
		}
	}
}
