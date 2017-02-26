// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.12, p. 74: write a function that reports whether two strings are
// anagrams of each other, that is, they contain the same letters in a
// different order.

// Use the "equal" map comparison function at p. 96.
package main

import (
	"fmt"
	"os"
)

type charMap map[rune]int

func main() {
	if equal(countChars(os.Args[1]), countChars(os.Args[2])) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func equal(x, y charMap) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func countChars(s string) charMap {
	counts := make(charMap)
	for _, r := range s {
		counts[r]++
	}
	return counts
}
