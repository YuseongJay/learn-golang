package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 2}
	v3 = Vertex{}
	v4 = Vertex{y: 3}
	p  = &Vertex{4, 9}
)

func main() {
	fmt.Println(v1, v2, v3, v4, *p)
}
