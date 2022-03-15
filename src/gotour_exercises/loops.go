package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	step := 1.0

	// Step size limited to a millionth
	for math.Abs(step) > 10e-6 {
		step = (z*z - x) / (2 * z)
		z -= step
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
