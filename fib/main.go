package main

import (
	"fib/lib"
	"fib/lib/add"
	"fmt"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	lib.Fib(c, quit)
	fmt.Println(add.Add(10, 30))
}
