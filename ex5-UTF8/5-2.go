package main

import "fmt"

func main() {
	x := 15
	max, min := 10, 0
	if x > max {
		fmt.Println("x max min", x, max, min)
		max += x
	} else {
		min = x
		fmt.Println("x max min", x, max, min)
	}
}
