package main

import "fmt"

func main() {
	//declare slice
	var slice []int
	fmt.Println(slice)

	// declare and initialization variable for slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// create slice from array
	var array = [4]int{4, 5, 6, 7}
	// from array[1] -> to array[3-1=2] <=> array[2]
	slice2 := array[1:3]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:] // from arr[2] -> end
	fmt.Println(slice4)

	slice5 := array[:3] // from start -> arr[3-1=2]
	fmt.Println(slice5)

	// create slice from other slice
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[3:]
	fmt.Println(slice8)

	// slice is => REFERENCE TYPE
	var array1 = [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]
	slice9[0] = 777
	fmt.Println(slice9)
	fmt.Println(array1)

	// length and capacity in slice
	// len: number element of slice -> ["FRANCE"] -> length 1
	// cap: number element of underlying array from start location when create slice -> ["FRANCE", "CANADA", "CHINA"] -> 3
	// capacity: capacity bộ nhớ mở rộng để chuẩn bị sẵn cho việc thêm phần tử vào slice
	countries := [...]string{"VN", "USA", "FRANCE", "CANADA", "CHINA"}
	slice10 := countries[2:3]
	fmt.Println(slice10)
	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))
	fmt.Println()

	// make copy, append
	// 1. make
	slide11 := make([]int, 2, 5)
	fmt.Println(slide11)
	fmt.Println(len(slide11))
	fmt.Println(cap(slide11))

	slide12 := make([]int, 2)
	fmt.Println(slide12)
	fmt.Println(len(slide12))
	fmt.Println(cap(slide12))
	fmt.Println()

	// 2. append
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)
	fmt.Println()

	// 4. copy
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)
	copy(dest, src)
	fmt.Println(dest)

	number := copy(dest, src)
	fmt.Println(number)
	fmt.Println()

	// delete element in slice
	src = append(src[:1], src[2:]...) // -> slice - slice = append(slice1, slice2...)
	fmt.Println(src)

	/*
		https://github.com/golang/go/wiki/SliceTricks
	*/
}
