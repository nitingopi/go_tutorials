package main

import (
	"fmt"
)

func panicRecover(){
	defer fmt.Println("deferred call 1")
	defer func() {
		fmt.Println("deferred call 2")
		if e:= recover(); e != nil { 
			fmt.Println("Recover with: ", e)
		}

	}()
	panic("Just panicking for the sake of example") //
	fmt.Println("This will be never be called")

}

func main() {
	fmt.Println("Starting to panic")
	panicRecover()
	fmt.Println("Program regains control after panic recovery")
}
