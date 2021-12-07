package main

import (
	"fmt"
)

func main() {
	for i, ch := range "korea 한국" {
		fmt.Printf("%d:%q", i, ch)
	}

	s := "korea 한국"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d:%q", i, s[i])
	}
}