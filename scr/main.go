package main

import "fmt"

func main() {
	// types we can store in go
	/*
		boolean

		integers
		floating point
		complex numbers

		text types
	*/
	// boolean data
	var n bool = true            // true
	fmt.Printf("%v, %T\n", n, n) // true, bool
	// by default all variables = 0
	var l bool                   // false
	fmt.Printf("%v, %T\n", l, l) // false, bool
	// integers variation
	/*
		int8	-128 - 127
		int16   -32 768 - 32 767
		int32   -2 147 483 648 - 2 147 483 647
		int64   far more

		also there is uintegers
		uint8    0 - 255
		uint16   0 - 65 535
		uint32   0 - 4 294 967 295
	*/

	// YOU CANT ADD TWO DIFFERENT TYPES OF INTEGERS!

	// binary operations & | ^ &^
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	// bit shifting
	a = 8               // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	// complex number

	var k complex64 = 1 + 2i                 // 1 + 2i
	fmt.Printf("%v, %T\n", k, k)             // (1+2i), complex64
	fmt.Printf("%v, %T\n", real(k), real(k)) // 1, float32
	fmt.Printf("%v, %T\n", imag(k), imag(k)) // 2, float32
	var i complex64 = complex(5, 12)         // 5 + 12i
	fmt.Printf("%v, %T\n", i, i)             // (5+12i), complex64

	// Text types
	s := "imma string"                                 // imma string
	fmt.Printf("%v, %T\n", s, s)                       // imma string, string
	fmt.Printf("%v, %T\n", s[0], s[0])                 // 105, uint8
	fmt.Printf("%v, %T\n", string(s[0]), string(s[0])) // i, string

	// adding strings
	s2 := "Imma string 2"
	twoStrings := s + s2
	fmt.Printf("%v, %T", twoStrings, twoStrings)
}
