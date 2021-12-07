package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 7, 11, 13}
	for i, v := range data {
		fmt.Println(i, v)
	}
}
