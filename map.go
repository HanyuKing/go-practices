package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["key"] = Vertex2{
		40.68433, -74.39967,
	}
	obj := m["key"]
	obj.Lat = float64(1)
	fmt.Println(m["key"])
}
