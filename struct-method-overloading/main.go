package main

import . "fmt"

type Point struct {
	X, Y int
}

func (p *Point) say() {
	Println("Say Point")
	return
}

func (l *Line) say() {
	Println("Say Line")
	return
}

type Line struct {
	Point
	name string
	id   int
}

func main() {
	var p Line
	p.say()
	p.Point.say()
}
