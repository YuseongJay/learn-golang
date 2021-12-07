package main

import "fmt"

func main() {
	a := 1
	a--
	b := 1.5
	b--
	c := 1 + 2i
	c--
	a--
	d := a
	a--
	e := a
	a--
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
