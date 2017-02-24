// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.6, p. 62: supersampling is a technique to reduce the effect of
// pixellation by computing the color value at several points within each
// pixel and taking the average. The simplest method is to divide each pixel
// into four "subpixels". Implement it.
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
		// Mandelbrot set.
		xmin, xmax, ymin, ymax = -2.1, +1.1, -1.4, +1.4
		// Antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.4865, -1.4705, -0.007, +0.007
		// Second antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.483885, -1.483805, -3.5e-5, +3.5e-5
		width, height = 2240, 1960
	)

	img := image.NewPaletted(image.Rect(0, 0, width/2, height/2), palette.Plan9)
	for py := 0; py < height; py += 2 {
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px += 2 {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := float64(px+1)/width*(xmax-xmin) + xmin
			colorIdx := uint8((mandelbrot(complex(x0, y0))+
				mandelbrot(complex(x0, y1))+
				mandelbrot(complex(x1, y0))+
				mandelbrot(complex(x1, y1)))/4.0 + 0.5)
			img.SetColorIndex(px/2, py/2, colorIdx)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) float64 {
	const iterations = 255

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2.0 {
			return float64(n)
		}
	}
	return 0.0
}
