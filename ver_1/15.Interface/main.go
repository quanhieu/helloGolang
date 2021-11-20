package main

import "fmt"

// Interface
type Animal interface {
	speak()
}

// Multiple interfaces
type Movement interface {
	move()
}

type Dog struct{}

// Embed interface
type NextAnimal interface {
	Movement
	Animal
}

func (d Dog) speak() {
	fmt.Println("Wan wan")
}

func (d Dog) move() {
	fmt.Println("Dog run on 4 legs")
}

// Empty interface
func goOut(i interface{}) {
	fmt.Println(i)
}

type data struct {
	index int
}

// pointer receiver interface
type Cat struct{}

func (d *Cat) speak() {
	fmt.Println("Nya Nya")
}

func (d *Cat) move() {
	fmt.Println("Cat jump on 4 legs")
}

func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()
	fmt.Println("=============")

	dog := Dog{}
	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()
	fmt.Println("=============")

	var nex NextAnimal = dog
	nex.speak()
	nex.move()
	fmt.Println("=============")

	goOut(10)
	goOut(12.13)
	fmt.Println()

	d := data{123}
	goOut(d)
	fmt.Println("=============")

	cat := Cat{}

	var next NextAnimal = &cat

	next.speak()
	next.move()
	fmt.Println("=============")

}
