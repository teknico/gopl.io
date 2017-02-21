// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on tempconv.
package weightconv

import "fmt"

type Kilograms float64
type Pounds float64

const (
	LbKgRatio = 0.45359237
	KgLbRatio = 1 / LbKgRatio
)

func (kg Kilograms) String() string { return fmt.Sprintf("%.2fkg", kg) }
func (lb Pounds) String() string    { return fmt.Sprintf("%.2flb", lb) }

// KgToLb converts a distance in Kilograms to Pounds.
func KgToLb(kg Kilograms) Pounds { return Pounds(kg * KgLbRatio) }

// LbToKg converts a distance in Pounds to Kilograms.
func LbToKg(lb Pounds) Kilograms { return Kilograms(lb * LbKgRatio) }
