package main

import "strings"

func main() {
	s := "Korea Seoul"

	println(strings.Contains(s, "abc"))
	println(strings.ContainsAny(s, "abc"))
	println(strings.Count(s, "ou"))
	println(strings.HasPrefix(s, "Ko"))
	println(strings.HasSuffix(s, "rea"))
	println(strings.Index(s, "ea"))
	println(strings.IndexAny(s, "abc"))
	println(strings.LastIndex(s, "eo"))
	println(strings.LastIndexAny(s, "abc"))
}
