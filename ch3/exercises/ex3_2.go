// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.2, p. 60: experiment with visualizations of other functions
// from the math package. Can you produce an egg box, moguls, or a saddle?
// Based on surface.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)

	// Egg box.
	xyrange = 20.0         // axis ranges (-xyrange..+xyrange)
	zscale  = height * 0.1 // pixels per z unit

	// Mogul.
	// xyrange = 3.0          // axis ranges (-xyrange..+xyrange)
	// zscale  = height * 0.3 // pixels per z unit

	// Saddle
	// xyrange = 2.5          // axis ranges (-xyrange..+xyrange)
	// zscale  = height * 0.1 // pixels per z unit
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	// Egg box.
	return math.Sin(x) * math.Sin(y)
	// Mogul.
	// return -x*x*x + x*x + x - 0.5
	// Saddle.
	// return x*x - y*y
}
