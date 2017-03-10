// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.9, p. 99: Write a program "wordfreq" to report the frequency of
// each word in an input text file. Call input.Split(bufio.ScanWords) before
// the first call to Scan to break the input into words instead of lines.
// Based on charcount.

// $ echo "Chattanooga Choo Choo won't you choo choo me home" | go run ex4_9.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ex. 4.9: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Word\tcount\n")
	for w, n := range words {
		fmt.Printf("%q\t%d\n", w, n)
	}
}
