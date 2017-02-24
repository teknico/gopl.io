// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.7, p. 62: another simple fractal uses Newton's method to find
// complex solutions to a function such as z⁴-1 = 0. Shade each starting point
// by the number of iterations required to get close to one of the four roots.
// Color each point by the root it approaches.
// Based on ex3.5 solution.
package main

import (
	"image"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		// Newton set.
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		// Zoom the center in.
		// xmin, ymin, xmax, ymax = -0.5, -0.5, +0.5, +0.5
		// Zoom in more.
		// xmin, ymin, xmax, ymax = -0.1, -0.1, +0.1, +0.1
		width, height = 1024, 1024
	)

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents a complex value.
			img.SetColorIndex(px, py, newton(complex(x, y)))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) uint8 {
	const iterations = 36
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return contrast * i
		}
	}
	return 0
}
