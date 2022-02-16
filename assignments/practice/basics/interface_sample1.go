package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}
	// fmt.Println(r.area())
	// fmt.Println(c.area())

	shapes := []shape{r, c}
	for _, s := range shapes {
		// fmt.Println(s.area())
		fmt.Println(getArea(s))
	}
}
