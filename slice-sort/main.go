package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	people = append(people, struct {
		Name string
		Age  int
	}{"엘리", 26})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)

	people = append(people, struct {
		Name string
		Age  int
	}{"고퍼", 29})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	a := "고리"
	b := "고퍼"
	fmt.Println(a, b, a < b)
}
