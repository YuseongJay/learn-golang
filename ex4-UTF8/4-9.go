package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Korea Seoul"

	fmt.Println(strings.Fields(s))
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.SplitAfter(s, " "))
	println(strings.Join([]string{"test", "result"}, "[]"))
	println(strings.Repeat(s, 2))
}
