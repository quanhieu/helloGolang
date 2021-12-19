package main

import (
	"fmt"
)

func main() {
	// function below DEFER will run at the end of program
	// defer is a stack -> LIFO
	fmt.Println("1 ")
	defer fmt.Println("2 ")
	fmt.Println("3 ")
	fmt.Println("4 ")

	a := 12
	defer fmt.Printf("\n %v a \n", a)
	a = 24
	fmt.Println("a ", a)

	// panic
	panicker()

	defer fmt.Println("DEFER ")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Make a panic")
}

func panicker() {
	fmt.Println("\nCreate panicker")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
			// throw panic to main
			// panic(err)
		}
	}()
	panic("Make a panicker \n")
	fmt.Println("Hello ")
}
