package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon ", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var p1 Pentagon = 100
	p1.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()
	o.Perimeter()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	// obj.Perimeter() can't do it
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
