package main

import "fmt"

func max(a, b int) int {
	//return a > b ? a : b
	var c int
	if a > b {
		c = a
	} else {
		c = b
	}

	return c
}

func main() {
	fmt.Println(max(3, 4))
}
