package main

import "fmt"

/*
Variadic function
	1. create variadic function
	2. pass a slice to variadic function
	3. change slice
*/

// append(slice [] Type, element ...Type) []Type
// append([]Tpe, args, arg2, argsN)

func addItem(item int, list ...int) {
	// 100, 200, 300, 400 -> int[] {100, 200, 300, 400}
	// []int {1,2,3,4} -> int[] {[]int {1,2,3,4}}

	// list = append(list, item, 11, 22, 33, 44, 55, 66)
	list = append(list, item)

	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 100, 200, 300, 400)
	fmt.Println()

	// implement slice in variadic function
	var list = []int{1, 2, 3, 4}
	addItem(100, list...)
	fmt.Println()

	// slice reference type
	change(list...)
	fmt.Println(list)

}
