package main

import (
	"fmt"
)

func main() {
	a := "korea"[:]
	b := "korea"[1:3]
	c := "korea"[:2]
	d := "korea"[2:]

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))
	fmt.Println(c, len(c))
	fmt.Println(d, len(d))
}
