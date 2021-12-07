package main

import (
	"runtime"
	"time"
)

func say(s string, n int) {
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		println(n, s)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 100; i++ {
		go say("world", i)
		say("hello", i)
		go say("inner", i)
	}
}
