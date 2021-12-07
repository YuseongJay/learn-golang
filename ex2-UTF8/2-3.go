package main

import "fmt"

func add(a, b int) int {
	return a + b
}

var a, b int = 1, 2

func main() {
	result := add(a, b)
	fmt.Printf("Result=%d = add(%d + %d)", result, a, b)
}
