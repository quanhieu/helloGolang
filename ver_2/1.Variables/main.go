package main

import (
	"fmt"
	"strconv"
)

var (
	n int     = 15
	m float64 = 20.2
)
var String string = "Export global scope"

func main() {
	var i int
	i = 11

	var j int = 12

	k := 13.5

	fmt.Println("n global scope ", n)
	var n int = 100
	fmt.Println("n shadow ", n)
	fmt.Println("m ", m)
	fmt.Println("string ", String)

	fmt.Println("i ", i)
	fmt.Println("j ", j)
	// %v -> variable | %T -> type
	fmt.Printf("%v , %T \n", k, k)

	var convertNumber int = int(k)
	fmt.Printf("%v , %T \n", convertNumber, convertNumber)

	var convertString string = strconv.Itoa(i)
	fmt.Printf("%v , %T \n", convertString, convertString)
}

/*
	1. Define variable and syntax
	2. Global and block scope
	3. Shadow
	4. Declared and not used
	5. Export global scope
	6. Naming convention -> camelcase
	7. Convert type
*/
