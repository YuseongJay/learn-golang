package main

import (
	. "fmt"
)

type Vertex struct {
	X, Y int
}

func (Vert *Vertex) mul() int {
	return Vert.X * Vert.Y
}

func (Vert *Vertex) mula(scal int) int {
	Vert.X *= scal
	Vert.Y *= scal
	return Vert.X * Vert.Y
}

func (Vert Vertex) mulb(scal int) int {
	Vert.X *= scal
	Vert.Y *= scal
	return Vert.X * Vert.Y
}

func (v Vertex) add(scal int) int {
	v.X += scal
	v.Y += scal
	return v.X + v.Y
}

func main() {
	vert := Vertex{10, 20}
	go Println(vert.mul())
	vert.X = 5
	vert.Y = 10
	go Println(vert.mul())

	vert = Vertex{10, 20}
	Println("mula()", vert.mula(10))
	Println("mulb()", vert.mulb(10))

	Println("add()", vert.add(10))
}
