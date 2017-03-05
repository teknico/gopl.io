// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.1, p. 84: Write a function that counts the number of bits that
// are different in two SHA256 hashes. (See PopCount from Section 2.6.2.)

package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	HexSize = sha256.Size * 2
)

// pc[i] is the population count of i.
var pc [256]int8

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + int8(i&1)
	}
}

// parseHexByte converts a two-rune string with two hex bytes to uint8.
func parseHexByte(s string) uint8 {
	v, err := strconv.ParseUint(s, 16, 8)
	if err != nil {
		log.Fatalln(s, "is not a hex byte.")
	}
	return uint8(v)
}

// numDiffBits returns the number of different bits of two SHA256 hashes.
func numDiffBits(h1, h2 string) int {
	numDiff := 0
	var hv1, hv2 uint8
	for l, r := 0, 2; l < HexSize; l, r = l+2, r+2 {
		hv1, hv2 = parseHexByte(h1[l:r]), parseHexByte(h2[l:r])
		numDiff += int(pc[hv1^hv2])
	}
	return numDiff
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Pass two SHA256 hashes as arguments.")
	}
	if len(os.Args[1]) != HexSize || len(os.Args[2]) != HexSize {
		log.Fatalln("Each hash needs to have", HexSize, "hex digits.")
	}
	fmt.Println(numDiffBits(os.Args[1], os.Args[2]), "different bits")
}
