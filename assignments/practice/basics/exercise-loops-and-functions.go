package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.4
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
