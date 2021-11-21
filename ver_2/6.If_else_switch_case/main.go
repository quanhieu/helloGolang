package main

import (
	"fmt"
)

func main() {
	// If - Else
	m := map[string]int{
		"Hieu": 19,
		"Tom":  21,
	}

	if age, isExist := m["Hieu"]; isExist {
		fmt.Println("age", age)
	}

	number := 90
	guess := 80

	if guess < 10 || !false || guess > 100 {
		fmt.Println("1")
	}

	if guess >= 10 && guess <= 100 {
		if guess == number {
			fmt.Println("guess = number")
		} else if guess < number {
			fmt.Println("guess nho hon number")
		}

		if guess > number {
			fmt.Println("guess lon hon number")
		}
	}

	if false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	// Switch - Case
	numberTwo := 3
	switch numberTwo {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Not others")
	}

	fmt.Println("")
	switch {
	case numberTwo <= 3:
		fmt.Println(">= 3")
		fallthrough
	case numberTwo <= 5:
		fmt.Println("<= 5")
		fmt.Println("33333")
		break
		fmt.Println("x3")
	default:
		fmt.Println("Not others")
	}

	fmt.Println("")
	var i interface{} = 1.5

	switch i.(type) {
	case int:
		fmt.Println("Type int")
	case float64:
		fmt.Println("Type float64")
	default:
		fmt.Println("Type other")
	}
}

/*

 */
