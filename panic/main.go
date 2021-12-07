package main

import "os"

func fin(f *os.File) {
	println("close")
	f.Close()
}

func openFile(fn string) {
	f, err := os.Open(fn)
	defer fin(f)
	if err != nil {
		panic(err)
	}
}

func main() {
	openFile("Invalid.txt")
	println("Done")
}
