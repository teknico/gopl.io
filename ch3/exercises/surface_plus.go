// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Experimental version including changes from ex. 3.1, 3.3 and 3.4, plus
// instrumentation reporting and "fixing" invalid function values.
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
		height, width := parseForm(r)
		w.Header().Set("Content-Type", "image/svg+xml")
		writeGraph(w, height, width)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseForm(r *http.Request) (height, width int) {
	height, width = 320, 600 // canvas size in pixels
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return height, width
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
	return height, width
}

func writeGraph(out io.Writer, height, width int) {
	max := 0.0
	min := 256.0
	zWeight := math.Sqrt(float64(height)*float64(width)) / 181.816
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"height='%d' width='%d'>\n", height, width)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, height, width)
			bx, by, bz := corner(i, j, height, width)
			cx, cy, cz := corner(i, j+1, height, width)
			dx, dy, dz := corner(i+1, j+1, height, width)
			zAvg := (az+bz+cz+dz)/zWeight + 46.0
			max, min = math.Max(max, zAvg), math.Min(min, zAvg)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'"+
				" style='fill: #%02x00%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, int(zAvg), 255-int(zAvg))
		}
	}
	// log.Printf("Max: %g, min: %g", max, min)
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j, height, width int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	xyscale := float64(width) / 2.0 / xyrange // pixels per x or y unit
	sx := float64(width)/2.0 + (x-y)*cos30*xyscale
	sz := z * float64(height) * 0.4
	sy := float64(height)/2.0 + (x+y)*sin30*xyscale - sz
	return sx, sy, sz
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	val := math.Sin(r) / r
	if math.IsInf(val, 0) || math.IsNaN(val) {
		log.Printf("Invalid function value at (%g, %g): %g\n", x, y, val)
		return 1.0
	}
	return val
}
