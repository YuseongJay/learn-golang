package main

import (
	"fmt"
)

func main() {
	imap := map[int]string{
		0: "North",
		1: "East",
		2: "West",
		3: "South",
	}
	fmt.Println(imap)
	fmt.Println(imap[1], imap[2])
}
