package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("Invalid.txt")
	println("Done")
}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r, "recovered")
		}
	}()

	_, err := os.Open(fn)
	if err != nil {
		fmt.Println("panic routine")
		panic(err)
	}
}
