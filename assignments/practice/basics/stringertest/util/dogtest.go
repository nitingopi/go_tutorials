package util

import "fmt"

type Dog struct {
	name  string
	breed string
}

func (d Dog) String() string {
	return fmt.Sprintf("My name is %s, I'm a %s", d.name, d.breed)
}
