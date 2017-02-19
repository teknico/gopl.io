// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.4, p. 13: modify dup2 to print the names of all files
// in which each duplicated line occurs.
// Based on dup2.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type lineStats struct {
	count     int
	filenames map[string]bool
}

func main() {
	counts := make(map[string]lineStats)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, stats := range counts {
		if stats.count > 1 {
			var fnames []string
			for fname := range stats.filenames {
				fnames = append(fnames, fname)
			}
			sort.Strings(fnames)
			fmt.Printf("%d\t%v\t%q\n", stats.count, fnames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]lineStats) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		// NOTE: ignoring potential errors from input.Err()
		ls, is_there := counts[line]
		if !is_there {
			ls = lineStats{0, make(map[string]bool)}
		}
		ls.count++
		ls.filenames[f.Name()] = true
		counts[line] = ls
	}
}
