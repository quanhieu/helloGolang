package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("i ", i)
	}

	fmt.Println("\n ")
	j, k := 1, 0
	for j <= 10 {
		if j%2 == 0 {
			fmt.Println("j ", j)
			fmt.Println("k ", k)
		}
		j++
		k++
	}
	// l := 1
	// for ;l <= 10; {
	// 	if l%2 == 0 {
	// 		fmt.Println("l ", l)
	// 	}
	// 	l++
	// }

	fmt.Println("\n ")
	m := 1
	for {
		{
			if m%2 != 0 {
				if m == 3 {
					m++
					continue
				}
				fmt.Println("m ", m)
			}
			m++
			if m > 10 {
				break
			}
		}
	}

	fmt.Println("\n ")
Loop:
	for i := 1; i <= 5; i++ {
		for n := 1; n <= 5; n++ {
			if i > 1 {
				break Loop
			}
			fmt.Println("i * n =", i*n)
		}
	}

	fmt.Println("\n ")
	array := [3]int{2, 3, 4}
	for i := 0; i < len(array); i++ {
		fmt.Println(i, "->", array[i])
	}

	fmt.Println("\n ")
	array2 := []int{2, 3, 4, 5, 6}
	for index, value := range array2 {
		fmt.Println(index, value)
	}

	fmt.Println("\n ")
	mapper := map[string]int{
		"Hieu": 20,
		"Tom":  22,
		"Mike": 25,
	}
	for index, value := range mapper {
		fmt.Println(index, value)
	}

	fmt.Println("\n ")
	s := "Hellow World"
	for index, value := range s {
		fmt.Println(index, string(value))
	}
	fmt.Println("\n ")
	for _, value := range s {
		fmt.Println(string(value))
	}
	fmt.Println("\n ")
	for index := range s {
		fmt.Println(index)
	}
}

/*

 */
