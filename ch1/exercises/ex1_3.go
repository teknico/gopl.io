// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.3, p. 8: experiment to measure the difference in running time
// between our potentially inefficient versions and the one that uses
// strings.Join.
// Based on echo2 and echo3.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	num_strs := []int{2, 3, 4, 5, 6, 8, 10, 30, 100, 1000, 10000, 100000}
	for _, num_str := range num_strs {
		fmt.Printf("%d strings:\n", num_str)
		var strs = make([]string, num_str, num_str)

		// Concatenation.
		start := time.Now()
		res := strs[0]
		for _, str := range strs[1:] {
			res += " " + str
		}
		concat_elapsed := time.Since(start).Seconds()
		fmt.Printf("\tconcat: %gs elapsed\n", concat_elapsed)

		// Join.
		start = time.Now()
		strings.Join(strs, " ")
		join_elapsed := time.Since(start).Seconds()
		fmt.Printf("\tjoin: %gs elapsed\n", join_elapsed)

		fmt.Printf("\tratio: %f\n", concat_elapsed/join_elapsed)
	}
}
