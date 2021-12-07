package main

import . "fmt"

func main() {
	f1 := func() {
		Println("Anonymous function")
	}
	f2 := func(a string) {
		Println(a)
	}
	f3 := func(a, b int) int {
		return a + b
	}

	f1()
	f2("Parameter function")
	Println(f3(5, 10))

	Println(&f1)
	Println(&f2)
	Println(&f3)
}
