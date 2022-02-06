package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Fibonacci example 1: %d \n\n", fib(10))

	number := 100
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan int, number)

	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}

	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < number; j++ {
		fmt.Printf("Fibonacci example 2: %d \n", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
