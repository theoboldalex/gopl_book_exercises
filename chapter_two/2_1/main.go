package main

import (
	"fmt"
	"tempconv/tempconv"
)

func main() {
	// test conversion from celsius to kelvin
	fmt.Printf("20°C: %v\n", tempconv.CToK(20)) // 295.15
	fmt.Printf("12°C: %v\n", tempconv.CToK(12)) // 285.15
	fmt.Printf("1°C: %v\n", tempconv.CToK(1))   // 274.15

	// test conversion from Kelivin to Celsius
	fmt.Printf("20K: %v\n", tempconv.KToC(20)) // -253.15°C
	fmt.Printf("12K: %v\n", tempconv.KToC(12)) // -261.15°C
	fmt.Printf("1K: %v\n", tempconv.KToC(1))   // -272.15°C
}
