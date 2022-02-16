package main

import "fmt"

type employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

type Emp struct {
	name  string
	basic int
}

func (e Emp) PrintName() {
	fmt.Println("Name:", e.name)
}

func (e Emp) PrintSalary(tax int) int {
	salary := e.basic + tax
	fmt.Println("Salary:", salary)
	return salary
}

func main() {
	// var e employee
	e := Emp{name: "John", basic: 10000}
	e.PrintName()
	e.PrintSalary(5000)
}
