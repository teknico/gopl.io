// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.9, p. 63: write a web server that renders fractals and writes
// the image data to the client. Allow the client to specify the x, y, and
// zoom values as parameters to the HTTP request.
// Example: http://localhost:8000/?x=-0.7501&y=0.0125&zoom=3000
// Based on ex3.5 solution.
package main

import (
	"image"
	"image/color/palette"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		x, y, zoom := parseForm(r)
		generateImage(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseForm(r *http.Request) (x, y, zoom float64) {
	x, y, zoom = 0.0, 0.0, 2.0 // defaults
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	req_x, is_there := r.Form["x"]
	if is_there {
		if num_x, err := strconv.ParseFloat(req_x[0], 64); err != nil {
			log.Printf("Invalid x: %q", req_x[0])
		} else {
			x = num_x
		}
	}
	req_y, is_there := r.Form["y"]
	if is_there {
		if num_y, err := strconv.ParseFloat(req_y[0], 64); err != nil {
			log.Printf("Invalid y: %q", req_y[0])
		} else {
			y = num_y
		}
	}
	req_zoom, is_there := r.Form["zoom"]
	if is_there {
		if num_zoom, err := strconv.ParseFloat(req_zoom[0], 64); err != nil {
			log.Printf("Invalid zoom: %q", req_zoom[0])
		} else {
			if num_zoom <= 0 {
				log.Printf("Zoom should be greater than zero: %q", num_zoom)
			} else {
				zoom = 1.0 / num_zoom
			}
		}
	}
	return x, y, zoom
}

func generateImage(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 800, 800
	)
	xmin, xmax := x-zoom, x+zoom
	ymin, ymax := y-zoom, y+zoom
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
	png.Encode(out, img) // NOTE: ignoring errors
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
