package main

import (
	"fmt"
)

func main() {
	m := map[string]float64{
		"pi": 3.1416,
		"e":  2.71828,
	}
	fmt.Println(m)
	for key, value := range m {
		fmt.Println(key, value)
	}
}
