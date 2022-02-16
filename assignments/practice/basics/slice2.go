package main

import "fmt"

func main() {
	var names = [6]string{"John", "Paul", "George", "Ringo", "Stuart", "Brian"}
	fmt.Println(names)
	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	names2 := []string{"John", "Paul", "George", "Maeve"}
	fmt.Println(names2)

	for _, v := range names2 {
		fmt.Println(v)
	}
}
