package main

import . "fmt"

type Point struct {
	X, Y int
}

func (p *Point) mula(scal int) int {
	p.X *= scal
	p.Y *= scal
	return p.X * p.Y
}

type Line struct {
	Point
	name string
	id   int
}

func main() {
	var p Line
	p.Point.X = 2
	p.Point.Y = 3

	Println(p.Point.mula(10))
	Println(p.mula(10))
}
