package main

import (
	"fmt"
	"math"
)

type Virtex struct {
	X, Y float64
}

func (v Virtex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Virtex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Virtex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
