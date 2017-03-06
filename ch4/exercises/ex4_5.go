// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.5, p. 93: Write an in-place function to eliminate adjacent
// duplicates in a []string slice.
// Based on nonempty.

package main

import "fmt"

// nodup returns a slice without adjacent duplicates.
// The underlying array is modified during the call.
func nodup(strings []string) []string {
	tip := 0
	for i, s := range strings {
		if s != strings[tip] {
			tip++
			if tip != i {
				strings[tip] = s
			}
		}
	}
	return strings[:tip+1]
}

func main() {
	data := []string{"a", "b", "b", "c", "d", "d", "e", "e"}
	fmt.Printf("%q\n", nodup(data)) // `["a" "b" "c" "d" "e"]`
	fmt.Printf("%q\n", data)        // `["a" "b" "c" "d" "e" "d" "e" "e"]`
}
