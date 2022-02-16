package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("deferred function 1, value of i is ", i) // will be called when returned from this function
	i++
	fmt.Println("The current value of i is ", i)
}