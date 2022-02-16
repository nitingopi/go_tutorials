package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 0
	count := 0
	return func() int {
		if count == 0 {
			first = 0
			count += 1
		} else if count == 1 {
			second = 1
			count += 1
		} else {
			first, second = second, first+second
		}

		return second
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
