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
}
