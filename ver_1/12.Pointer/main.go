package main

import "fmt"

func main() {
	a := 100

	var pointer *int

	// zero value
	fmt.Println("zero value ", pointer)

	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T ", pointer)
	fmt.Println()

	b := 123
	p2 := new(int) // <=> var p2 *int
	p2 = &b

	fmt.Println(p2)
	fmt.Printf("%T ", p2)
	fmt.Println()

	pointer2 := new(int)
	fmt.Println("pointer2 ", pointer2)
	fmt.Println()

	// access value of pointer
	var pointer3 *int
	c := 100
	pointer3 = &c
	fmt.Println(c)

	*pointer3 = 999 // <=> a := 999
	fmt.Println(pointer3)
	fmt.Println(c)
	fmt.Println()

	// demo pointer -> array
	array := [3]int{1, 2, 3}
	var pointer4 *[3]int
	pointer4 = &array
	fmt.Println(pointer4)
	fmt.Printf("Type %T ", pointer4)
	fmt.Println()

	// access pointer through function
	d := 30
	var pointer5 *int = &d
	applyPointer(pointer5)
	fmt.Println(d)

}

func applyPointer(pointer *int) {
	*pointer = 777
}
