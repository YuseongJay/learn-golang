package main

import "fmt"

func main() {
	data := []int{2, 3, 5}
	fmt.Printf("data[:3] len(data) %d cap(data) %d \n", len(data), cap(data))
	data = append(data, 7, 11)
	fmt.Printf("data[:3] len(data) %d cap(data) %d \n", len(data), cap(data))

	str := append([]byte("hello "), "world"...)
	fmt.Println(str)
}
