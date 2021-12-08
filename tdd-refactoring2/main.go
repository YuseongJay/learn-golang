package main

import (
	"fmt"
)

func main() {
	test()
}

func assertEqual(testcase string, expected, o interface{}) bool {
	if expected != o {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := calculate(op, a, b)
	return assertEqual(testcase, expected, o)
}

type Testcase struct {
	testcase, op   string
	a, b, expected int
}

var testcases []Testcase = []Testcase{
	{
		testcase: "test1",
		op:       "+",
		a:        1,
		b:        2,
		expected: 3,
	},
	{
		testcase: "test2",
		op:       "+",
		a:        1,
		b:        -2,
		expected: -1,
	},
	{
		testcase: "test3",
		op:       "-",
		a:        1,
		b:        -2,
		expected: 3,
	},
	{
		testcase: "test4",
		op:       "-",
		a:        1,
		b:        2,
		expected: -1,
	},
	{
		testcase: "test5",
		op:       "*",
		a:        1,
		b:        2,
		expected: 2,
	},
	{
		testcase: "test6",
		op:       "*",
		a:        3,
		b:        7,
		expected: 21,
	},
	{
		testcase: "test7",
		op:       "*",
		a:        -2,
		b:        5,
		expected: -10,
	},
	{
		testcase: "test8",
		op:       "*",
		a:        -3,
		b:        -1,
		expected: 3,
	},
	{
		testcase: "test9",
		op:       "*",
		a:        3,
		b:        -4,
		expected: -12,
	},
	{
		testcase: "test10",
		op:       "/",
		a:        3,
		b:        1,
		expected: 3,
	},
	{
		testcase: "test11",
		op:       "/",
		a:        10,
		b:        2,
		expected: 5,
	},
	{
		testcase: "test12",
		op:       "/",
		a:        7,
		b:        2,
		expected: 3,
	},
	{
		testcase: "test13",
		op:       "/",
		a:        -5,
		b:        3,
		expected: -1,
	},
	{
		testcase: "test14",
		op:       "/",
		a:        11,
		b:        11,
		expected: 1,
	},
}

func test() {
	for i := 0; i < len(testcases); i++ {
		testcase := testcases[i].testcase
		op := testcases[i].op
		a := testcases[i].a
		b := testcases[i].b
		expected := testcases[i].expected

		if !testCalculate(testcase, op, a, b, expected) {
			return
		}
	}

	fmt.Println("Success!")
}

func calculate(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Unknown operator: " + op)
	}
}
