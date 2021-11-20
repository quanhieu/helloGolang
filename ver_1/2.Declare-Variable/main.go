package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	// Type Inference
	var email = "ryan@gmail.com"
	fmt.Println(email, "\n")

	// Declare multi variables
	// 1. Same variable
	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)

	var a1, b1 = 10, 11
	fmt.Println(a1, b1)

	var a2, b2 = "aaa", 333
	fmt.Println(a2, b2, "\n")

	// 2. Different variable
	var (
		name    string
		address string
		age     int
	)

	name = "Robin"
	address = "Vietnam"
	age = 25

	fmt.Println(name, address, age)

	var (
		name1    string = "Batman"
		address1 string = "USA"
		age1     int    = 36
	)

	fmt.Println(name1, address1, age1, "\n")

	language := "Golang"
	fmt.Println(language)
}
