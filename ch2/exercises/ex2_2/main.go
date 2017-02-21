// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.2, p. 44: write a general-purpose unit-conversion program
// analogous to cf that reads numbers from its command-line arguments or from
// the standard input if there are no arguments, and converts each number into
// units like temperature in Celsius and Fahrenheit, length in feet and meters,
// weight in pounds and kilograms, and the like.
// Based on cf.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../ex2_1/tempconv"
	"./lengthconv"
	"./weightconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			args = append(args, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
			continue
		}
		fmt.Println(t)
		tc, tf := tempconv.Celsius(t), tempconv.Fahrenheit(t)
		fmt.Printf("\t%s = %s\t|\t%s = %s\n",
			tc, tempconv.CToF(tc), tf, tempconv.FToC(tf))
		lm, lft := lengthconv.Meters(t), lengthconv.Feet(t)
		fmt.Printf("\t%s = %s\t|\t%s = %s\n",
			lm, lengthconv.MToFt(lm), lft, lengthconv.FtToM(lft))
		wkg, wlb := weightconv.Kilograms(t), weightconv.Pounds(t)
		fmt.Printf("\t%s = %s\t|\t%s = %s\n",
			wkg, weightconv.KgToLb(wkg), wlb, weightconv.LbToKg(wlb))
	}
}
