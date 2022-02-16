package main

import (
	"fmt"
)

var langs map[string]string

func addItem(k string, v string) {
	if _, ok := langs[k]; ok {
		fmt.Println("key already exists")
		return
	}
	langs[k] = v
	fmt.Println("key is inserted")

}

func updateItem(k string, v string) {
	langs[k] = v
	fmt.Println("key value updated")
}

func getById(k string) (v string) {
	if _, ok := langs[k]; ok {
		return langs[k]
	}
	return ""
}

func getAll(langs map[string]string) map[string]string {
	return langs
}

func deleteItem(k string) {
	delete(langs, k)
}

func init() {

	langs = make(map[string]string)
}

func main() {
	fmt.Println("Hello world")
	addItem("TCS", "Tata Consultancy Services")
	addItem("CTS", "Cognizant Technology Services")
	addItem("IBM", "International Business Machines")
	addItem("IBM", "International Business Machines")
	deleteItem("TCS")
	deleteItem("TCS")
	fmt.Println(getAll(langs))
	updateItem("TCS", "Tata Group")

	fmt.Println(langs)
	fmt.Println("Hello world")

}
