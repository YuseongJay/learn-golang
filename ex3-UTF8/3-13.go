package main

var x string = "Hello World"

func f() {
	println(x)
	var x string = "f()function"
	println(x)
}

func main() {
	var x string = "main function"
	println(x)
	f()
}
