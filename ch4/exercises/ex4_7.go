// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.7, p. 93: Modify reverse to reverse the characters of
// a []byte slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?
// Based on rev.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := []byte(string("a\u4e16b\u4e16c"))
	reverse(data)
	fmt.Printf("%q\n", string(data)) // "c世b世a"
}

func reverse(b []byte) {
	// First reverse each code point larger than one byte.
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		if size > 1 {
			for j, k := i, i+size-1; j < k; j, k = j+1, k-1 {
				b[j], b[k] = b[k], b[j]
			}
		}
		i += size
	}
	// Then reverse the whole thing.
	for j, k := 0, len(b)-1; j < k; j, k = j+1, k-1 {
		b[j], b[k] = b[k], b[j]
	}
}
