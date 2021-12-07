package main

import . "fmt"

func main() {
	var c []string
	var d []string
	c = []string{"sun", "mon", "tue", "wed"}
	d = []string{"aa", "aaa", "aaaaa"}
	Println(c, c[0], c[1], c[0][0])
	Println(d)
	Println(d[0])
	Println("d[0][0],d[0][1]", d[0][0], d[0][1])
}
