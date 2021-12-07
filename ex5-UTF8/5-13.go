package main

import "fmt"

func main() {
	const s = "서울 Korea"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}
}
