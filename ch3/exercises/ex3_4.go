// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.4, p. 60: following the approach of the Lissajous example in
// Section 1.7, construct a web server that computes surfaces and writes SVG
// data to the client. The server must set the Content-Type header like this:
//     w.Header().Set("Content-Type", "image/svg+xml")
// Allow the client to specify values like height, width, and color as HTTP
// request parameters.
// Based on surface.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		height, width, color := parseForm(r)
		w.Header().Set("Content-Type", "image/svg+xml")
		writeGraph(w, height, width, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseForm(r *http.Request) (height, width int, color string) {
	height, width = 320, 600 // canvas size in pixels
	color = "white"          // polygons color
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return height, width, color
	}
	req_height, is_there := r.Form["height"]
	if is_there {
		if num_height, err := strconv.Atoi(req_height[0]); err != nil {
			log.Printf("Invalid height: %q", req_height[0])
		} else {
			height = num_height
		}
	}
	req_width, is_there := r.Form["width"]
	if is_there {
		if num_width, err := strconv.Atoi(req_width[0]); err != nil {
			log.Printf("Invalid width: %q", req_width[0])
		} else {
			width = num_width
		}
	}
	// To use hex colors in the URL, escape the hash as %23:
	// http://localhost:8000/?color=%2300ff00
	req_color, is_there := r.Form["color"]
	if is_there {
		color = req_color[0]
	}
	return height, width, color
}

func writeGraph(out io.Writer, height, width int, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"height='%d' width='%d'>\n", color, height, width)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width)
			bx, by := corner(i, j, height, width)
			cx, cy := corner(i, j+1, height, width)
			dx, dy := corner(i+1, j+1, height, width)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j, height, width int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	xyscale := float64(width) / 2.0 / xyrange // pixels per x or y unit
	sx := float64(width)/2.0 + (x-y)*cos30*xyscale
	sy := float64(height)/2.0 + (x+y)*sin30*xyscale - z*float64(height)*0.4
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
