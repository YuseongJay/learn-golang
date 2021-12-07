package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()

	data := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("data=", data)
	fmt.Printf("len(data) %d cap(data) %d \n", len(data), cap(data))
	data = data[1:4]
	fmt.Println("data[1:4]=", data)
	fmt.Printf("data[1:4] len(data) %d cap(data) %d \n", len(data), cap(data))
	data = data[1:2]
	fmt.Println("data[1:2]=", data)
	fmt.Printf("data[1:2] len(data) %d cap(data) %d \n", len(data), cap(data))
	data = data[:3]
	fmt.Println("data[:3]=", data)
	fmt.Printf("data[:3] len(data) %d cap(data) %d \n", len(data), cap(data))
}
