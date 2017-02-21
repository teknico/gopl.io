// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Changes by Nicola 'tekNico' Larosa <gopl.io@teknico.net>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0.0
	BoilingC      Celsius    = 100.0
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32.0
	BoilingF      Fahrenheit = 212.0
	AbsoluteZeroK Kelvin     = 0.0
	FreezingK     Kelvin     = 273.15
	BoilingK      Kelvin     = 373.15
	CFRatio                  = 1.8
	FCRatio                  = 1 / CFRatio
	KFRatio                  = 1.8
	FKRatio                  = 1 / KFRatio
	CKRatio                  = 1.0
	KCRatio                  = 1.0
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }
