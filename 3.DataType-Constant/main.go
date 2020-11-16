package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Bool: true/false
	var myBool bool = true
	// String
	var myString string = "Hello"
	// int
	var myInt int = 123

	fmt.Println(myBool, myString, myInt)

	// int 8, 16, 32, 64
	// 1. Range
	fmt.Println(math.MinInt8)       // -128
	fmt.Println(math.MaxInt8, "\n") // 127

	fmt.Println(math.MinInt16)       // -32768
	fmt.Println(math.MaxInt16, "\n") // 32767

	fmt.Println(math.MinInt32)       // -2147483648
	fmt.Println(math.MaxInt32, "\n") // 2147483647

	fmt.Println(math.MinInt64)       // -9223372036854775808
	fmt.Println(math.MaxInt64, "\n") // 9223372036854775807

	// 2. Bits
	fmt.Println(bits.OnesCount8(math.MaxUint8), "\n")   // 8
	fmt.Println(bits.OnesCount16(math.MaxUint16), "\n") // 16
	fmt.Println(bits.OnesCount32(math.MaxUint32), "\n") // 32
	fmt.Println(bits.OnesCount64(math.MaxUint64), "\n") // 64

	// uint - Positive integer
	var myUInt uint = 10
	fmt.Println(myUInt, "\n")

	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("byte is alias for %T", myByte)
	fmt.Println("\n")

	var maxByte byte = 255
	fmt.Println(maxByte, " is byte value", math.MaxUint8, "\n")

	var val byte = 'E'
	fmt.Println("Unicode val is %X", val, "\n")

	// uintptr
	/* Save pointer address
	key - value
	key = uintpre
	value = object (*point)

	key = (*point) -> uintptr
	map.put(key, value)
	*/

	// float
	var myFloat float64 = 10.01
	fmt.Println(myFloat, "\n")

	// complex z = a + bi
	var z1 complex64 = 10 + 1i
	var z2 complex64 = 10 + 1i
	fmt.Println(z1, z2, z1+z2, "\n")

	// Rune
	var newString string = "Nhẫn"
	for i := 0; i < len(newString); i++ {
		fmt.Printf("%c", newString[i])
	}
	fmt.Println("\n")

	runes := []rune(newString)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println("\n")

	var myRune rune = '☘'
	fmt.Printf("fmt golang format %T - %c - %U - %d", myRune, myRune, myRune, myRune)
	fmt.Println("\n")

	// convert type value
	var a int = 1
	var b float64 = 2.1
	fmt.Println(float64(a)+b, "\n")

	// constants cannot be declared using :=syntax
	const PI = 3.14159
	fmt.Println(PI)
}
