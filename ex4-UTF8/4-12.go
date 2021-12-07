package main

import "fmt"

const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38   // 2**127 * (2**24-1)/2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45  // 1/2**(127-1+23)
	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023*(2**53-1)/2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1/2**(1023-1+52)
)

func main() {
	fmt.Println("MaxFloat32", MaxFloat32)
	fmt.Println("SmallestNonzeroFloat32", SmallestNonzeroFloat32)
	fmt.Println("MaxFloat64", MaxFloat64)
	fmt.Println("SmallestNonzeroFloat64", SmallestNonzeroFloat64)
}
