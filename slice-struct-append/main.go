package main

import (
	"fmt"
)

type test struct {
	Title string
	Body  string
}

func main() {
	a := test{Title: "a", Body: "a body"}
	b := test{Title: "b", Body: "b body"}

	var arr []test
	arr = append(arr, a)
	arr = append(arr, b)

	fmt.Println(arr)
}
