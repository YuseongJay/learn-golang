package main

import "fmt"

func main() {
	var a uint8 = 3
	b := ^a
	fmt.Printf("%08b", b)
}
