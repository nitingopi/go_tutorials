package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {

	n := make(map[string]Vertex)
	n["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m)
	fmt.Println(n)

}
