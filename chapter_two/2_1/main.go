package main

import (
	"fmt"
	"tempconv/tempconv"
)

func main() {
	// test conversion from Celsius to Kelvin
	fmt.Printf("20°C -> K: %v\n", tempconv.CToK(20)) // 295.15
	fmt.Printf("12°C -> K: %v\n", tempconv.CToK(12)) // 285.15
	fmt.Printf("1°C -> K: %v\n", tempconv.CToK(1))   // 274.15

	// test conversion from Kelvin to Celsius
	fmt.Printf("20K -> °C: %v\n", tempconv.KToC(20)) // -253.15°C
	fmt.Printf("12K -> °C: %v\n", tempconv.KToC(12)) // -261.15°C
	fmt.Printf("1K -> °C: %v\n", tempconv.KToC(1))   // -272.15°C

	// test conversion from Fahrenheit to Kelvin
	fmt.Printf("20°F -> K: %v\n", tempconv.FToK(20)) // -266.48°C
	fmt.Printf("12°F -> K: %v\n", tempconv.FToK(12)) // -262.04°C
	fmt.Printf("1°F -> K: %v\n", tempconv.FToK(1))   // -255.93°C

	// test conversion from Kelvin to Fahrenheit
	fmt.Printf("20K -> °F: %v\n", tempconv.KToF(20)) // -423.67°F
	fmt.Printf("12K -> °F: %v\n", tempconv.KToF(12)) // -438.07°F
	fmt.Printf("1K -> °F: %v\n", tempconv.KToF(1))   // -457.87°F
}
