package main

import "fmt"

func f1() {
	defer fmt.Println("f1 defer")
	fmt.Println("f1 program")
}

func main() {
	defer fmt.Println("main defer")
	f1()
	fmt.Println("main program")
}
