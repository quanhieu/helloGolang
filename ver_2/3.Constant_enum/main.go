package main

import (
	"fmt"
)

const a = 42
const (
	// _ = iota + 5
	_ = 1 << iota * 10
	redColor
	yellow
	blue
	black = "abc"
)

func main() {
	fmt.Printf("%v , %T\n", a, a)
	const a int16 = 12
	fmt.Printf("%v , %T\n\n", a, a)

	fmt.Printf("%v , %T\n", redColor, redColor)
	fmt.Printf("%v , %T\n", yellow, yellow)
	fmt.Printf("%v , %T\n", blue, blue)
	fmt.Printf("%v , %T\n", black, black)
}

/*
	1. Describe constant enum
	2. Define
	3. Chareactor of
	4. Type of Variable
	5. Auto define variable enum by iota
*/
