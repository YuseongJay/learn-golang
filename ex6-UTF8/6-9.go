package main

import . "fmt"

func main() {
	var c []float32
	var d [][]float32
	c = []float32{0.1234567890, 1234567890.0}
	d = [][]float32{{1.1, 11.1, 111.1},
		{2.2, 22.2, 222.2},
		{3, 4, 5}}
	Println(c[0], c[1])
	Println(d[0], d[1][1])
}
