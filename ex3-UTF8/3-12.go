package main

func f() {
	println(x)
}

func main() {
	println(x)
	f()
}

var x string = "Hello World"
