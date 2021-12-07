package main

import "fmt"

func main() {
	x := 15
	max := 10
	if x > max {
		fmt.Println("x max", x, max)
		max += x
	}

	fmt.Println("x max", x, max)
}
