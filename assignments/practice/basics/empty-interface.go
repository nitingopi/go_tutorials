package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
