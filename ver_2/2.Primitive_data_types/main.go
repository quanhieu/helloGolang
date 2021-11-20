package main

import (
	"fmt"
)

func main() {
	// Boolean
	var a bool
	fmt.Printf("%v , %T\n", a, a)
	a = 1 != 2
	fmt.Printf("%v , %T\n", a, a)
	a = 1 == 2
	fmt.Printf("%v , %T\n \n", a, a)

	// Numerics - Integer
	var b int8 = 9
	var c int8 = 3
	fmt.Printf("%v , %T\n", b, b)
	fmt.Printf("%v , %T\n", b+c, b+c)
	fmt.Printf("%v , %T\n", b-c, b-c)
	fmt.Printf("%v , %T\n", b*c, b*c)
	fmt.Printf("%v , %T\n", b/c, b/c)
	fmt.Printf("%v , %T\n \n", b%c, b%c)

	fmt.Printf("%v , %T\n", b&c, b&c)
	fmt.Printf("%v , %T\n", b|c, b|c)
	fmt.Printf("%v , %T\n", b^c, b^c)
	fmt.Printf("%v , %T\n", b&^c, b&^c)
	fmt.Printf("%v , %T\n", b<<3, b<<3)
	fmt.Printf("%v , %T\n\n", b>>3, b>>3)

	// Numerics - Float
	var d float64 = 3.14
	d = 13.7e72
	d = 2.1e72
	fmt.Printf("%v , %T\n\n", d, d)

	// Numerics - Complex
	var e complex64 = 1 + 2i
	var f complex128 = complex(5, 12)
	fmt.Printf("%v , %T\n", e, e)
	fmt.Printf("%v , %T\n", real(e), real(e))
	fmt.Printf("%v , %T\n", imag(e), imag(e))
	fmt.Printf("%v , %T\n", f, f)
	fmt.Printf("%v , %T\n", real(f), real(f))
	fmt.Printf("%v , %T\n\n", imag(f), imag(f))

	// String
	var g string = "Hello"
	var h string = "world"
	fmt.Printf("%v , %T\n\n", g+" "+h, g+h)

	// Byte
	var i string = "Hello golang"
	var j []byte = []byte(i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n\n", string(j), j)

	// Rune
	var k rune = 'ă'
	fmt.Printf("%v , %T\n", k, k)
	fmt.Printf("%v , %T\n", string(k), k)
}

/*
	Primitive Date Types
	1. Boolean
	2. Numerics
			Integer
					signed integer	int8 int16 int32 int64
					unsigned integer	uint8 uint16 uint32 uint64
			Float		float32 float64
			Complex		complex64 complex128
	3. Text
			String
			Byte 	-> UTF-8
			Rune	-> UTF-32
*/
