package main

import (
	"fmt"
)

var imap map[int]string

func main() {
	imap = make(map[int]string)
	imap[0] = "North"
	imap[1] = "East"
	imap[2] = "West"
	imap[3] = "South"

	fmt.Println(imap)
	fmt.Println(imap[1], imap[2])
}
