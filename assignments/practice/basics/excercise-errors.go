package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (f ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Math: sqrt of negative number %g", float64(f))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	var z float64 = 1.4
	z -= (z*z - x) / (2 * z)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
