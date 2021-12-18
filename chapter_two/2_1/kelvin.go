// Add types, constants and functions to tempconv for
// processing temps in the kelvin scale
package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

var (
	AbsoluteZero Kelvin = 273.15
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c) + AbsoluteZero }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - AbsoluteZero) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9) + AbsoluteZero }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-AbsoluteZero)*9/5 + 32) }
