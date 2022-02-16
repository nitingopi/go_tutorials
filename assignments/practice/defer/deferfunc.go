package main 

import (
	"fmt"
	"io/ioutil" 
	"os"
)


func Readfile(filename string) ([] byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close() // will execute when returned from this function
	return ioutil.ReadAll(f)
}

func main() {
	f,_ := Readfile("test.txt")
	fmt.Println("%s",f)
	fmt.Println(string(f))
}