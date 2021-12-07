package main

import "fmt"

func main() {
	s := "서울 Korea"
	for i, a := range s {
		fmt.Printf("%#U start at position %d\n", a, i)
	}
}
