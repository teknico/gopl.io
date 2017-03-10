// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.8, p. 99: Modify charcount to count letters, digits, and so on
// in their Unicode categories, using functions like unicode.IsLetter.
// Based on charcount.

// $ echo "\b5Ὂg̀9! ℃ᾭG" | go run ex4_8.go
// Rune counts
// Control: 1
// Digit: 2
// Graphic: 12
// Letter: 5
// Lower: 2
// Mark: 1
// Number: 2
// Print: 12
// Punct: 2
// Space: 1
// Symbol: 1
// Title: 1
// Upper: 2

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	Control = iota
	Digit
	Graphic
	Letter
	Lower
	Mark
	Number
	Print
	Punct
	Space
	Symbol
	Title
	Upper
	_NumRuneTypes
)

func classifyChar(r rune, cnt []int) {
	if unicode.IsControl(r) {
		cnt[Control]++
	}
	if unicode.IsDigit(r) {
		cnt[Digit]++
	}
	if unicode.IsGraphic(r) {
		cnt[Graphic]++
	}
	if unicode.IsLetter(r) {
		cnt[Letter]++
	}
	if unicode.IsLower(r) {
		cnt[Lower]++
	}
	if unicode.IsMark(r) {
		cnt[Mark]++
	}
	if unicode.IsNumber(r) {
		cnt[Number]++
	}
	if unicode.IsPrint(r) {
		cnt[Print]++
	}
	if unicode.IsPunct(r) {
		cnt[Punct]++
	}
	if unicode.IsSpace(r) {
		cnt[Space]++
	}
	if unicode.IsSymbol(r) {
		cnt[Symbol]++
	}
	if unicode.IsTitle(r) {
		cnt[Title]++
	}
	if unicode.IsUpper(r) {
		cnt[Upper]++
	}
}

func main() {
	counts := make([]int, _NumRuneTypes)
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex. 4.8: %v\n", err)
			os.Exit(1)
		}
		classifyChar(r, counts)
	}
	fmt.Println("Rune counts")
	fmt.Println("Control:", counts[Control])
	fmt.Println("Digit:", counts[Digit])
	fmt.Println("Graphic:", counts[Graphic])
	fmt.Println("Letter:", counts[Letter])
	fmt.Println("Lower:", counts[Lower])
	fmt.Println("Mark:", counts[Mark])
	fmt.Println("Number:", counts[Number])
	fmt.Println("Print:", counts[Print])
	fmt.Println("Punct:", counts[Punct])
	fmt.Println("Space:", counts[Space])
	fmt.Println("Symbol:", counts[Symbol])
	fmt.Println("Title:", counts[Title])
	fmt.Println("Upper:", counts[Upper])
}
