package main

import "fmt"

// none-struct
type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s%s", s, str)
}

type Student struct {
	name string
}

// Define method
/*
	func (t Type) methodName(params) return { // body code }
	(t Type) => Receiver
	1. Value Receiver -> not effect
	2. Pointer Receiver -> effect
*/
func (s Student) getName() string {
	return s.name
}

// 1. Value Receiver
func (s Student) changeName() {
	fmt.Printf("p2 = %p", &s)
	s.name = "Robin"
}

// 2. Pointer Receiver
func (s *Student) changeName2() {
	fmt.Printf("p4 = %p", s)
	s.name = "Hood"
}

func main() {
	student := Student{"Ryan"}
	name := student.getName()
	fmt.Println(name)
	fmt.Println()

	// 1. Value Receiver
	fmt.Printf("p1 = %p", &student)
	fmt.Println()

	student.changeName()
	fmt.Println()
	fmt.Println(student.name)
	fmt.Println()

	// 2. Pointer Receiver
	fmt.Printf("p3 = %p", &student)
	fmt.Println()

	student.changeName2()
	fmt.Println()
	fmt.Println(student.name)
	fmt.Println()

	// none-struct
	s1 := String("a")
	newStr := s1.append("b")
	fmt.Println(newStr)
	fmt.Println()
}
