package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func g1() {
	fmt.Println("g1")
}

func g2() {
	fmt.Println("g2")
	wg.Done()
}

func g3() {
	fmt.Println("g3")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// Goroutines
	go g1()

	countGoroutines := runtime.NumGoroutine()
	fmt.Println(countGoroutines)

	time.Sleep(time.Second)
	fmt.Println()

	// Synchronized goroutines
	/*
		login 1
			go g2()
			go g3()

		logic 2
	*/
	fmt.Println("Begin")

	wg.Add(2)
	go g2()
	go g3()

	wg.Wait()
	fmt.Println("End")

	
}