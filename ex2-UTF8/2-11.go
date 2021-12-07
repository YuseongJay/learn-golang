package main

import "fmt"

func main() {
	a := 1
	b := &a
	fmt.Println(b)

	c := new(int)
	*c = 1
	fmt.Println(*c)
}
