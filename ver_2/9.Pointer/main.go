package main

import (
	"fmt"
)

func main() {
	var a int = 12
	var b *int = &a
	fmt.Println("1 ", a, b)
	fmt.Println("1 ", a, *b)
	a = 24
	fmt.Println("2 ", a, b)
	fmt.Println("2 ", a, *b)
	*b = 32
	fmt.Println("3 ", a, *b)

	// array
	c := [3]int{1, 2, 3}
	d := c
	fmt.Println("\narray ", c, d)
	c[1] = 9
	fmt.Println("array ", c, d)

	// slice
	e := []int{1, 2, 3}
	f := e
	fmt.Println("\nslice ", e, f)
	e[1] = 9
	e[2] = 9
	fmt.Println("slice ", e, f)
	fmt.Println("slice ", &e[2], &f[2])

	// map
	var h = map[string]string{"Hieu": "Male", "Kate": "FEMALE"}
	i := h
	fmt.Println("\nmap ", h, i)
	h["Kate"] = "Female"
	fmt.Println("map ", h, i)

	// struct
	var g *myStruct
	fmt.Println("\nstruct ", g)
	g = new(myStruct)
	fmt.Println("struct ", g)
	(*g).number = 20
	fmt.Println("struct ", (*g).number)
}

type myStruct struct {
	number int
}
