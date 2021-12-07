package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 7, 11, 13}
	var s = make([]int, 3)
	_ = copy(s, data)
	fmt.Printf("s len(s) %d cap(s) %d n%d \n", len(s), cap(s), s)
	_ = copy(s, s[:2])
	fmt.Printf("s len(s) %d cap(s) %d n%d \n", len(s), cap(s), s)

	s = s[:2]
	fmt.Printf("s len(s) %d cap(s) %d n%d \n", len(s), cap(s), s)
	var i = make([]int, 5)
	_ = copy(i, s[:2])
	fmt.Printf("i len(i) %d cap(i) %d n%d \n", len(i), cap(i), i)
}
