package main

import "fmt"

func main() {
	a := 112
	a >>= 3
	fmt.Printf("%08b", a)
}
