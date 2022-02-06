package main

import (
	"fmt"
)

func main() {
	example1()

	example2()
}

func example1() {
	randomNumbers := []int{}
	for i := 1; i <= 100000000; i++ {
		randomNumbers = append(randomNumbers, i)
	}
	sum := 0
	for i := 0; i < len(randomNumbers); i++ {
		sum += randomNumbers[i] * randomNumbers[i]
	}
	fmt.Printf("Total sum of square example1: %d \n", sum)
}

func example2() {
	randomNumbers := []int{}
	for i := 1; i <= 100000000; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	// generate pipeline
	inputChan := generatePipeline(randomNumbers)

	// Fan-out
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	// Fan-in
	c := fanIn(c1, c2, c3, c4)

	sum := 0
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("Total sum of square example2: %d \n", sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChannel ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChannel {
			for n := range c {
				in <- n
			}
		}
	}()
	return in
}
