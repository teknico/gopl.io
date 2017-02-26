// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.10, p. 74: write a non-recursive version of comma, using
// bytes.Buffer instead of string concatenation.
// Based on comma.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 123456 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	123,456
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	rem := n % 3
	if rem == 0 {
		rem = 3
	}
	out := s[:rem]
	for idx := rem; idx < n; idx += 3 {
		out += "," + s[idx:idx+3]
	}
	return out
}
