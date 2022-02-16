package main

import (
	"fmt"
	"modules/lib/mymath"
	"modules/lib/mystring"
)

func main() {

	sum := mymath.Add(3, 4)
	fmt.Println("sum", sum)
	reversedstring := mystring.Reverse("hello")
	fmt.Println("reversedstring", reversedstring)
}
