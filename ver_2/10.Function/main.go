package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	max := findMax(a, b)
	fmt.Println("1 ", max)

	maxSecond := findMaxSecond(a, b)
	fmt.Println("1.2 ", maxSecond)

	c := 30
	d := 40
	fmt.Println("2 ", pointerFindMax(&c, &d))

	maxPointer := reUsePointerFindMax(c, d)
	fmt.Println("3 ", *maxPointer)

	e := []int{1, 2, 3, 4, 5, 6}
	sum := findSum(e)
	fmt.Println("\n4 ", sum)

	sumSecond := findSumSecond("Heck", 1, 2, 3, 4, 5)
	fmt.Println("4.2 ", sumSecond)

	res, err := calculateDivide(4, 0)
	if err != nil {
		fmt.Println("\n5 ", err)
	}
	fmt.Println("5.2 ", res)

	func() {
		fmt.Println("\n6 Anonymous function ")
	}()

	func() {
		fmt.Println("6.2 Anonymous function - run parallel ")
		for i := 0; i <= 5; i++ {
			func(i int) {
				fmt.Println(i)
			}(i)
		}
	}()

	g := greeter{
		greeting: "Hello",
		name:     "Golang",
	}
	g.greet()
	fmt.Println("7 ", g.name)

	g.greetPointer()
	fmt.Println("7.2 ", g.name)
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxSecond(a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func pointerFindMax(a, b *int) int {
	*a = 100
	if *a > *b {
		return *a
	}
	return *b
}

func reUsePointerFindMax(a, b int) *int {
	a = 200
	if a > b {
		return &a
	}
	return &b
}

func findSum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func findSumSecond(s string, a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	fmt.Println("Hello ", s)
	return
}

func calculateDivide(a, b int) (int, error) {
	result := 0
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result = a / b
	return result, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println("\n7 ", g.greeting, g.name)
	g.name = ""
}

func (g *greeter) greetPointer() {
	fmt.Println("\n7 ", g.greeting, g.name)
	g.name = "Pointer"
}
