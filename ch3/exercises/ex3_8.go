// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.8, p. 63: rendering fractals at high zoom levels demands great
// arithmetic precision. Implement the same fractal using four different
// representations of numbers: complex64, complex128, big.Float, and big.Rat.
// How do they compare in performance and memory usage? At what zoom levels do
// rendering artifacts become visible?
// Based on ex3.5 solution.
//
// CPU times on an Intel Core i7-4750HQ, 2.00GHz, with Go 1.8:
// Complex64:  #1 0.082s, #2 0.12s, #3 0.22s, #4 0.22s.
// Complex128: #1 0.072s, #2 0.11s, #3 0.20s, #4 0.27s.
// big.Float:  #1    30s, #2   49s, #3  104s, #4  149s.
// big.Rat:    #1 aborted: 38s per line, would take hours to do all.
//                #2, #3 and #4 won't even finish one pixel (!).
package main

import (
	"image"
	"image/color/palette"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"os"
)

// big.Float mantissa precision bits (float64 precision is 53 bit).
const bfprec = 200

// Since Newton's Method doubles the number of correct digits at each
// iteration, we need at least log_2(bfprec) steps for big.Float.Sqrt.
const bfsteps = 8

func main() {
	const (
		width, height = 300, 300
		// #1: Mandelbrot set.
		xmin, xmax, ymin, ymax = -0.5 - 1.6, -0.5 + 1.6, -1.4, +1.4
		// #2: antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.4785 - 0.008, -1.4785 + 0.008, -0.007, +0.007
		// #3: second antenna zoom in.
		// xmin, xmax, ymin, ymax = -1.483845 - 4e-5, -1.483845 + 4e-5, -3.5e-5, +3.5e-5
		// #4: third zoom in.
		// xmin, xmax = -1.4840000361 - 1e-14, -1.4840000361 + 1e-14
		// ymin, ymax = -5.200001e-8 - 1e-14, -5.200001e-8 + 1e-14
	)

	// big.Float
	// mbf := NewMandelBigFloat()
	// big.Rat
	// mbr := NewMandelBigRat()

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
	for py := 0; py < height; py++ {
		// log.Printf("*** Line #%d", py)
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			// log.Printf("Line %d, pixel #%d", py, px)
			x := float64(px)/width*(xmax-xmin) + xmin

			// complex64
			// img.SetColorIndex(px, py, mandelbrotComplex64(x, y))
			// complex128
			img.SetColorIndex(px, py, mandelbrotComplex128(x, y))
			// big.Float
			// img.SetColorIndex(px, py, mbf.getPixel(x, y))
			// big.Rat
			// img.SetColorIndex(px, py, mbr.getPixel(x, y))
		}
	}
	log.Println("Finished, writing image...")
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotComplex64(x, y float64) uint8 {
	const iterations = 255

	z := complex64(complex(x, y))
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2.0 {
			return n
		}
	}
	return 0
}

func mandelbrotComplex128(x, y float64) uint8 {
	const iterations = 255

	z := complex(x, y)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2.0 {
			return n
		}
	}
	return 0
}

// Custom implementations of big.Float- and bit.Rat-based complex numbers.
// These types do not have a square root method, needed for computing the
// absolute value, so we have to implement that too.

type MandelBigFloat struct {
	xpx    *big.Float
	ypx    *big.Float
	xacc   *big.Float
	yacc   *big.Float
	xsq    *big.Float
	ysq    *big.Float
	cross  *big.Float
	sqr    *big.Float
	sqracc *big.Float
	sqrtmp *big.Float
	two    *big.Float
	half   *big.Float
}

func NewMandelBigFloat() *MandelBigFloat {
	mbf := new(MandelBigFloat)
	mbf.xpx = new(big.Float).SetPrec(bfprec)
	mbf.ypx = new(big.Float).SetPrec(bfprec)
	mbf.xacc = new(big.Float).SetPrec(bfprec)
	mbf.yacc = new(big.Float).SetPrec(bfprec)
	mbf.xsq = new(big.Float).SetPrec(bfprec)
	mbf.ysq = new(big.Float).SetPrec(bfprec)
	mbf.cross = new(big.Float).SetPrec(bfprec)
	mbf.sqr = new(big.Float).SetPrec(bfprec)
	mbf.sqracc = new(big.Float).SetPrec(bfprec)
	mbf.sqrtmp = new(big.Float).SetPrec(bfprec)
	mbf.two = new(big.Float).SetPrec(bfprec).SetFloat64(2.0)
	mbf.half = new(big.Float).SetPrec(bfprec).SetFloat64(0.5)
	return mbf
}

func (mbf *MandelBigFloat) getPixel(x, y float64) uint8 {
	const iterations = 255

	mbf.xpx.SetFloat64(x)
	mbf.ypx.SetFloat64(y)
	mbf.xacc.SetFloat64(0.0)
	mbf.yacc.SetFloat64(0.0)
	for n := uint8(0); n < iterations; n++ {
		// sq = acc*acc where (complex multiplication)
		// xsq = xacc*xacc - yacc*yacc
		mbf.xsq.Mul(mbf.xacc, mbf.xacc)
		mbf.ysq.Mul(mbf.yacc, mbf.yacc)
		mbf.xsq.Sub(mbf.xsq, mbf.ysq)
		// ysq = yacc*xacc + xacc*yacc
		mbf.cross.Mul(mbf.xacc, mbf.yacc)
		mbf.ysq.Add(mbf.cross, mbf.cross)
		// acc = sq + px
		mbf.xacc.Add(mbf.xsq, mbf.xpx)
		mbf.yacc.Add(mbf.ysq, mbf.ypx)
		// abs(acc) = sqrt(xacc*xacc + yacc*yacc)
		mbf.xsq.Mul(mbf.xacc, mbf.xacc)
		mbf.ysq.Mul(mbf.yacc, mbf.yacc)
		mbf.sqr.Add(mbf.xsq, mbf.ysq)
		// sqrt(sqr) -> sqracc (as x), sqrtmp as t, in Newton method.
		mbf.sqracc.SetFloat64(1.0) // Initial estimate.
		for i := 0; i <= bfsteps; i++ {
			mbf.sqrtmp.Quo(mbf.sqr, mbf.sqracc)    // t = sqr / x_n
			mbf.sqrtmp.Add(mbf.sqracc, mbf.sqrtmp) // t = x_n + (sqr / x_n)
			mbf.sqracc.Mul(mbf.half, mbf.sqrtmp)   // x_{n+1} = 0.5 * t
		}
		if mbf.sqracc.Cmp(mbf.two) == 1 {
			return n
		}
	}
	return 0
}

type MandelBigRat struct {
	xpx    *big.Rat
	ypx    *big.Rat
	xacc   *big.Rat
	yacc   *big.Rat
	xsq    *big.Rat
	ysq    *big.Rat
	cross  *big.Rat
	sqr    *big.Rat
	sqracc *big.Rat
	sqrtmp *big.Rat
	two    *big.Rat
	half   *big.Rat
}

func NewMandelBigRat() *MandelBigRat {
	mbr := new(MandelBigRat)
	mbr.xpx, mbr.ypx = new(big.Rat), new(big.Rat)
	mbr.xacc, mbr.yacc = new(big.Rat), new(big.Rat)
	mbr.xsq, mbr.ysq = new(big.Rat), new(big.Rat)
	mbr.cross, mbr.sqr = new(big.Rat), new(big.Rat)
	mbr.sqracc, mbr.sqrtmp = new(big.Rat), new(big.Rat)
	mbr.two = new(big.Rat).SetFloat64(2.0)
	mbr.half = new(big.Rat).SetFloat64(0.5)
	return mbr
}

func (mbr *MandelBigRat) getPixel(x, y float64) uint8 {
	const iterations = 255

	mbr.xpx.SetFloat64(x)
	mbr.ypx.SetFloat64(y)
	mbr.xacc.SetFloat64(0.0)
	mbr.yacc.SetFloat64(0.0)
	for n := uint8(0); n < iterations; n++ {
		// sq = acc*acc where (complex multiplication)
		// xsq = xacc*xacc - yacc*yacc
		mbr.xsq.Mul(mbr.xacc, mbr.xacc)
		mbr.ysq.Mul(mbr.yacc, mbr.yacc)
		mbr.xsq.Sub(mbr.xsq, mbr.ysq)
		// ysq = yacc*xacc + xacc*yacc
		mbr.cross.Mul(mbr.xacc, mbr.yacc)
		mbr.ysq.Add(mbr.cross, mbr.cross)
		// acc = sq + px
		mbr.xacc.Add(mbr.xsq, mbr.xpx)
		mbr.yacc.Add(mbr.ysq, mbr.ypx)
		// abs(acc) = sqrt(xacc*xacc + yacc*yacc)
		mbr.xsq.Mul(mbr.xacc, mbr.xacc)
		mbr.ysq.Mul(mbr.yacc, mbr.yacc)
		mbr.sqr.Add(mbr.xsq, mbr.ysq)
		// sqrt(sqr) -> sqracc (as x), sqrtmp as t, in Newton method.
		mbr.sqracc.SetFloat64(1.0) // Initial estimate.
		for i := 0; i <= bfsteps; i++ {
			mbr.sqrtmp.Quo(mbr.sqr, mbr.sqracc)    // t = sqr / x_n
			mbr.sqrtmp.Add(mbr.sqracc, mbr.sqrtmp) // t = x_n + (sqr / x_n)
			mbr.sqracc.Mul(mbr.half, mbr.sqrtmp)   // x_{n+1} = 0.5 * t
		}
		if mbr.sqracc.Cmp(mbr.two) == 1 {
			return n
		}
	}
	return 0
}
