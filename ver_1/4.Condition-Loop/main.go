package main

import "fmt"

func main() {
	// condition
	number := 10
	// if condition
	if number == 10 {
		fmt.Println("Number = 10")
	}

	if number < 10 {
		fmt.Println("Number < 10", "\n")
	} else {
		fmt.Println("Number >= 10", "\n")
	}

	// if statement; condition
	if a := 100; a > 100 {
		fmt.Println("Number > 10", "\n")
	} else {
		fmt.Println("Number <= 10", "\n")
	}

	// switch - case
	number1 := 3
	switch number1 {
	case 1:
		fmt.Println("Switch number 1", "\n")
	case 4, 200, 3:
		fmt.Println("Switch number 4", "\n")
	case 11:
		fmt.Println("Switch number 11", "\n")
	default:
		fmt.Println("Unknow", "\n")
	}

	switch {
	case number1 < 2:
		fmt.Println("Switch number < 2", "\n")
	case number1 > 2:
		fmt.Println("Switch number > 2", "\n")
	default:
		fmt.Println("Unknow2", "\n")
	}

	// FallThrough, break, goto
	switch number1 {
	case 1:
		fmt.Println("Switch number = 1")
		fallthrough
	case 2:
		fmt.Println("Switch number = 2")
		fallthrough
	case 3:
		fmt.Println("Switch number = 3")
		fallthrough
	case 4:
		fmt.Println("Switch number = 4")
		fallthrough
	case 5:
		fmt.Println("Switch number = 5")
		fallthrough
	case 6:
		if number == 3 {
			goto handleEqualNumber
		}
		fmt.Println("Break here")
	handleEqualNumber:
		fmt.Println("Handle for case = 3", "\n")
	default:
		fmt.Println("Unknow")
	}

	// Loops: for init; condition; post
	// break; continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			fmt.Println("Break loop at ", i)
			break
		}
	}
	fmt.Println("Out of Loop", "\n")

	// multiple condition
	for i, j := 1, 2; i < 7 && j < 10; i, j = i+1, j+2 {
		fmt.Println("i = ", i)
		fmt.Println("j = ", j)
	}
	fmt.Println("\n")

	// while
	j := 0
	for j < 10 {
		fmt.Println("while is ", j)
		j++

		if j > 9 {
			fmt.Println("\n")
		}
	}

	// infinite loop
	/*
		for {
			fmt.Println("Declare infinite loop")
		}
	*/
}
