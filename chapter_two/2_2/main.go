// Ex 2.2
// Write a general purpose unit conversion program that reads values from
// command line arguments (or stdin if there are none) and converts each
// number into the following units...
// celsius, fahrenheit, feet, meters, pounds and kilograms

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl/chapter_two/2_2/tempconv"
)

func main() {
	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Cannot convert argument to correct number format")
			os.Exit(1)
		}

		c = tempconv.Celsius(i)
		f = tempconv.CToF(c)
	}

	fmt.Printf("°C: %.1f°C\n°F: %.1f°F\n", c, f)
}
