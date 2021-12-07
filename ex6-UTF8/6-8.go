package main

import . "fmt"

func main() {
	var c [2]float32
	var d []float32
	c[0] = 0.1234567890
	c[1] = 1234567890.0
	d = []float32{1.1, 11.1, 111.1, 1111.1}
	Println(c[0], c[1])
	Println(d[0], d[1])
}
