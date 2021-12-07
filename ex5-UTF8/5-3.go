package main

import "fmt"

func main() {
	if x, max, min := 15, 10, 0; x > max {
		fmt.Println("x max min", x, max, min)
		max += x
	} else if x < min {
		min = x
		fmt.Println("x max min", x, max, min)
	} else {
		fmt.Println("x max min", x, max, min)
	}
}
