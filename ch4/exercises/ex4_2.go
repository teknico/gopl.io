// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.2, p. 84: Write a program that prints the SHA256 hash of its
// standard input by default but supports a command-line flag to print the
// SHA384 or SHA512 hash instead.
// Based on echo4 from ch. 2.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var hashName = flag.String("hash", "256", "one of '256', '384', '512'")

func main() {
	flag.Parse()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	switch *hashName {
	case "256":
		fmt.Printf("SHA256: %x\n", sha256.Sum256(data))
	case "384":
		fmt.Printf("SHA384: %x\n", sha512.Sum384(data))
	case "512":
		fmt.Printf("SHA512: %x\n", sha512.Sum512(data))
	default:
		log.Fatalln("Hash has to be one of '256', '384', '512'")
	}
}
