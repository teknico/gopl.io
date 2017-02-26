// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.13, p. 78: write const declarations for KB, MB, up through YB
// as compactly as you can.
package main

const (
	mult = 1000
	KB   = mult
	MB   = KB * mult
	GB   = MB * mult
	TB   = GB * mult
	PB   = TB * mult
	EB   = PB * mult
	ZB   = EB * mult
	YB   = ZB * mult
)

func main() {
}
