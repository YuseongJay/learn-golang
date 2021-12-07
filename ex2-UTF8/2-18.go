package main

import "fmt"

func main() {
	a := 7
	a <<= 2
	fmt.Printf("%08b", a)
}
