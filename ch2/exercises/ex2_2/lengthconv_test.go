// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run "go test lengthconv_test.go".
package lengthconv_test

import (
	"math"
	"testing"

	lc "./lengthconv"
)

const Epsilon = 1e-13

func TestFtToMToFt(t *testing.T) {
	var testdata = []struct {
		ft lc.Feet
		m  lc.Meters
	}{
		{lc.Feet(0), lc.Meters(0)},
		{lc.Feet(1), lc.Meters(lc.FtMRatio)},
		{lc.Feet(lc.MFtRatio), lc.Meters(1)},
	}

	for _, data := range testdata {
		mExp := lc.FtToM(data.ft)
		if math.Abs(float64(mExp-data.m)) > Epsilon {
			t.Errorf("FtToM(%q) = %v; want %v", data.ft, mExp, data.m)
		}
		ftExp := lc.MToFt(data.m)
		if math.Abs(float64(ftExp-data.ft)) > Epsilon {
			t.Errorf("MToFt(%q) = %v; want %v", data.m, ftExp, data.ft)
		}
	}
}
