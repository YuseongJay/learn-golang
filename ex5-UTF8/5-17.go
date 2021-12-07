package main

import (
	"fmt"
	"time"
)

func main() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
