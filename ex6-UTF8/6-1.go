package main

import . "fmt"

func main() {
	var c [2]string
	var d []string
	c[0] = "first"
	c[1] = "second"
	d = []string{"aa", "aaa", "aaaaa"}
	Println(c, c[0], c[1], c[0][0])
	Println(d)
	Println(d[0])
	Println("d[0][0],d[0][1]", d[0][0], d[0][1])
}
