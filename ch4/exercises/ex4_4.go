// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.4, p. 93: Write a version of rotate that operates
// in a single pass.
// Based on rev.

package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	rotate(s, 1)
	fmt.Println(s) // "[3 4 5 0 1 2]"
	rotate(s, -3)
	fmt.Println(s) // "[0 1 2 3 4 5]"
}

// rotate rotates a slice of ints by n places.
func rotate(s []int, n int) {
	if n < 0 {
		n = len(s) + n
	}
	for i, j := 0, n; i != j; {
		s[i], s[j] = s[j], s[i]
		i++
		j++
		if j == len(s) {
			j = n
		} else if i == n {
			n = j
		}
	}
}
