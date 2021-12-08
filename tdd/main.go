package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	test()
	elapsedTime := time.Since(startTime)
	fmt.Printf("runtime:%d \n", elapsedTime.Nanoseconds())
}

func assertEqual(expect, o interface{}) {
	if expect != o {
		fmt.Printf("Test Failed! expected %d output: %d", expect, o)
	}
}

func test() {
	var o interface{}
	o = calculate("+", 1, 2)
	assertEqual(3, o)
	o = calculate("+", 1, -2)
	assertEqual(-1, o)

	o = calculate("-", 1, -2)
	assertEqual(3, o)
	o = calculate("-", 1, 2)
	assertEqual(-1, o)

	o = calculate("*", 1, 2)
	assertEqual(2, o)
	o = calculate("*", 3, 7)
	assertEqual(21, o)
	o = calculate("*", -2, 5)
	assertEqual(-10, o)
	o = calculate("*", -3, -1)
	assertEqual(3, o)
	o = calculate("*", 3, -4)
	assertEqual(-12, o)

	o = calculate("/", 3, 1)
	assertEqual(3, o)
	o = calculate("/", 10, 2)
	assertEqual(5, o)
	o = calculate("/", 7, 2)
	assertEqual(3, o)
	o = calculate("/", -5, 3)
	assertEqual(-1, o)

	fmt.Println("Success!")
}

func calculate(op string, a, b int) int {
	var rst int
	if op == "+" {
		rst = a + b
	} else if op == "-" {
		rst = a - b
	} else if op == "*" {
		rst = a * b
	} else if op == "/" {
		rst = a / b
	}

	return rst
}
