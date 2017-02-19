// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.

// Exercise 1.12, p. 22: modify the Lissajous server to read parameter values
// from the URL. For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead
// of the default 5. Use the strconv.Atoi function to convert the string
// parameter into an integer. You can see its documentation with
// go doc strconv.Atoi.
// Based on lissajous.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x0, 0xff, 0x0, 0xff}, // green
}

const (
	bgIndex = 0 // first color in palette
	fgIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	cycles := 5 // number of complete x oscillator revolutions

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			} else {
				req_cycles, is_there := r.Form["cycles"]
				if is_there {
					if num_cycles, err := strconv.Atoi(req_cycles[0]); err != nil {
						log.Printf("Invalid cycles: %q", req_cycles[0])
					} else {
						cycles = num_cycles
					}
				}
			}
			lissajous(w, cycles)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				fgIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
