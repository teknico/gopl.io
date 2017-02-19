// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.7, p. 17: the function call io.Copy(dst, src) reads from src
// and writes to dst. Use it instead of ioutil.ReadAll to copy the response
// body to os.Stdout without requiring a buffer large enough to hold the
// entire stream. Be sure to check the error result of io.Copy.
// Based on fetch.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1.7: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1.7: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
