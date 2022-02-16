package main

import "fmt"

func getType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	default:
		fmt.Println("Unknown")			
	}
}

func main() {
	getType(25)
	getType("hello")
	getType(true)
}