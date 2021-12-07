package main

import (
	. "fmt"
)

func main() {
	adddr := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos, neg := adddr(), adddr()

	for i := 0; i < 10; i++ {
		Println(pos(i), neg(-2*i))
	}

	sum := 0
	addr := func(x int) int {
		sum += x
		return sum
	}

	aum := 0
	negg := func(x int) int {
		aum += x
		return aum
	}

	for i := 0; i < 10; i++ {
		Println(addr(i), negg(-2*i))
	}
}
