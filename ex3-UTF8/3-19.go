package main

import "fmt"

const (
	num0, num1 = iota, iota * 10
	num2, num3
	num4, num5
)

func main() {
	fmt.Println(num0, num1, num2, num3, num4, num5)
}
