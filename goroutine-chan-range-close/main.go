package main

func fibonacci(n int, c chan uint64) {
	var x, y uint64 = 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan uint64, 100)
	go fibonacci(cap(c), c)
	for i := range c {
		println(i)
	}
}
