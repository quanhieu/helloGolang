package main

import "fmt"

// func func_name (params) return_type { code }
func Hello() {
	fmt.Println("Hello")
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

// Multiple return values
func rectInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// Named return values
func rectInfoSecond(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	// sayHello("Ryan")
	result := greeting("Alice")
	fmt.Println(result, "\n")

	w, h, area := rectInfo(100, 300)
	fmt.Println("width = ", w)
	fmt.Println("height = ", h)
	fmt.Println("area = ", area, "\n")

	w1, h1, isSquare := rectInfoSecond(500, 500)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("width = ", w1)
		fmt.Println("height = ", h1)
	}

}
