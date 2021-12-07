package main

import "fmt"

func and() {
	a := 3
	b := 2
	c := a & b
	fmt.Printf("%08b\n", c)
}

func or() {
	a := 3
	b := 2
	c := a | b
	fmt.Printf("%08b\n", c)

}

func xor() {
	a := 3
	b := 2
	c := a ^ b
	fmt.Printf("%08b\n", c)
}

func andNot() {
	a := 255
	b := 68
	c := a &^ b
	fmt.Printf("%08b", c)
}

func main() {
	and()
	or()
	xor()
	andNot()
}
