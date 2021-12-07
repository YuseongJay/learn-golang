package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i, k := range m {
		fmt.Println(k, i)
	}
}
