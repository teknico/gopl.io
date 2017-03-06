// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.6, p. 93: Write an in-place function that squashes each run of
// adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte
// slice into a single ASCII space.
// Based on nonempty.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	tip := 0
	isCurSpace, isPrevSpace := false, false
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		isCurSpace = unicode.IsSpace(r)
		switch {
		case !isPrevSpace && !isCurSpace:
			tip += size
		case isPrevSpace && !isCurSpace:
			// Assume all spaces have size one, or we'd need prevSize too.
			tip += 1
			for j := 0; j < size; j++ {
				b[tip+j] = b[i+j]
			}
			tip += size
		case !isPrevSpace && isCurSpace:
			b[tip] = ' '
		}
		// case isPrevSpace && isCurSpace:
		//     Skip current space and keep scanning.
		isPrevSpace = isCurSpace
		i += size
	}
	return b[:tip]
}

func main() {
	data := []byte(string("\u4e16\t\n\v x\f\r\u0085\u00A0\u4e16"))
	fmt.Printf("%q\n", string(squashSpaces(data))) // "世 x 世"
}
