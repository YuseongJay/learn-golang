package main

import "strings"

func main() {
	s := "Korea Seoul"

	println(strings.Replace(s, "o", "_", 1))
	println(strings.ToUpper(s))
	println(strings.ToLower(s))
	println(strings.Title(s))
	println(strings.TrimSpace(s))
	println(strings.Trim(s, "Seoul"))
	println(strings.TrimLeft(s, "S"))
	println(strings.TrimRight(s, "S"))
	println(strings.TrimPrefix(s, "Korea"))
	println(strings.TrimSuffix(s, "Seoul"))
}
