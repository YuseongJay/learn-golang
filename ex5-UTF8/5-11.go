package main

import "fmt"

func main() {
	s := []string{"sum", "mon", "tue", "wed", "thu", "fri", "sat"}
	for i, a := range s {
		fmt.Println(i, a)
	}
}
