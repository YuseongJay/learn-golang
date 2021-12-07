package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("directory name:")
	var dir string
	fmt.Scanf("%s", &dir)
	os.Mkdir(dir, 0755)
	os.WriteFile(dir+"/main.go", []byte("package main\n\nfunc main() {\n\n}"), os.ModePerm)
}
