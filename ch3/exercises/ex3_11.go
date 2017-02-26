// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.11, p. 74: enhance comma so that it deals correctly with
// floating-point numbers and an optional sign.
// Based on comma.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 -1 0.1 -0.1 12 123 1234 -1234.56 123456 1234567890 -1234567890.123
// 	1
//  -1
//  0.1
//  -0.1
// 	12
// 	123
// 	1,234
//  -1234.56
// 	123,456
// 	1,234,567,890
//  -1234567890.123
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a number string.
func comma(s string) string {
	prefix := ""
	if strings.IndexAny(s, "+-") == 0 {
		prefix = string(s[0])
		s = s[1:]
	}
	dotIdx := strings.Index(s, ".")
	suffix := ""
	if dotIdx != -1 {
		suffix = s[dotIdx:]
		s = s[:dotIdx]
	}
	intPart := comma_int(s)
	return prefix + intPart + suffix
}

// comma_int inserts commas in a non-negative decimal integer string.
func comma_int(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma_int(s[:n-3]) + "," + s[n-3:]
}
