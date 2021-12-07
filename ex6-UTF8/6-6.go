package main

import . "fmt"

func main() {
	var c [2]int
	d := []int{11, 111, 1111}
	c[0] = 5
	c[1] = 55
	Println(c, c[0], c[1])
	Println(d)
	Println("d[0],d[1]", d[0], d[1])
}
