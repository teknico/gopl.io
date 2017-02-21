// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on tempconv.
package lengthconv

import "fmt"

type Feet float64
type Meters float64

const (
	FtMRatio = 0.3048
	MFtRatio = 1 / FtMRatio
)

func (ft Feet) String() string  { return fmt.Sprintf("%.2fft", ft) }
func (m Meters) String() string { return fmt.Sprintf("%.2fm", m) }

// FToM converts a distance in feet to meters.
func FtToM(ft Feet) Meters { return Meters(ft * FtMRatio) }

// MToF converts a distance in meters to feet.
func MToFt(m Meters) Feet { return Feet(m * MFtRatio) }
