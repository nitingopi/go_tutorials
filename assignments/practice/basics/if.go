package main 

import (
	"fmt"
)

func isEven(input int) {
	if input %2 == 0 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func main() {
	isEven(4)
}