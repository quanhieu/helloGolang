package main

import (
	"fmt"
	"time"

	"github.com/quanhieu/helloGolang/pattern/singleton"
)

func main() {
	example1()

	example2()
}

func example1() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()

	fmt.Println("AddOne s1 ", s1.AddOne())
	fmt.Println("AddOne s1 ", s1.AddOne())
	fmt.Println("AddOne s2 ", s2.AddOne())
	fmt.Println("AddOne s2 ", s2.AddOne())

	fmt.Printf("Singleton example1 1: %p \n", s1)
	fmt.Printf("Singleton example1 2: %p \n\n", s2)
}

func example2() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Singleton example2: %p \n", singleton.GetInstanceGoroutines())
		}()
	}

	time.Sleep(time.Second * 10)
}
