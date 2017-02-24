// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.5, p. 62: implement a full-color Mandelbrot set using
// the function image.NewRGBA and the type color.RGBA or color YCbCr.
// Based on mandelbrot.
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
		// Mandelbrot set.
		xmin, xmax, ymin, ymax = -2.1, +1.1, -1.4, +1.4
		// Antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.4865, -1.4705, -0.007, +0.007
		// Second antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.483885, -1.483805, -3.5e-5, +3.5e-5
		width, height = 1120, 980
	)

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.SetColorIndex(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) uint8 {
	const iterations = 255

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2.0 {
			return n
		}
	}
	return 0
}
