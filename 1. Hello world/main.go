package main

import "fmt"

func main() {
	fmt.Println("Hello golang")
	fmt.Println(Sum(3))
	fmt.Println("What the heck")
}

func Sum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}
