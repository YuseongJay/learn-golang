package main

import . "fmt"

func main() {
	func() {
		Println("Anonymous function")
	}()

	func(a string) {
		Println(a)
	}("Anonymous function")

	res := func(a, b int) int {
		return a + b
	}(1, 2)
	Println(res)
}
