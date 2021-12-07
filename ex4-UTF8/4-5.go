package main

import (
	"strings"
)

func main() {
	if "Korea" == "korea" {
		println("1Korea==korea")
	}
	if strings.EqualFold("Korea", "korea") {
		println("2Korea==korea")
	}
	if "Korea" > "korea" {
		println("korea grater then Korea")
	}
}
