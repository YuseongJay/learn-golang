package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 5.1
	b := int(a)
	c := uint16(0x10fe)
	d := int8(c)
	e := uint16(d)
	f := float64(1.9)
	g := int64(f)
	h := int64(-f)
	i := 1234567890
	j := float32(i)
	s1 := string(97)
	s2 := string(-1)
	s3 := strconv.Itoa(97)
	s4 := string([]byte{97, 230, 151, 165})
	s5 := []byte("a日")

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	fmt.Println(s1, s2, s3, s4, s5)
}
