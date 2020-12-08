package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza %d \n", pizza)
}

func timeTrack(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}

func main() {
	defer timeTrack(time.Now(), "Build Pizzas")

	number_of_evens := 3
	runtime.GOMAXPROCS(number_of_evens)
	wg.Add(number_of_evens)

	for i:= 0; i< number_of_evens; i++ {
		go createPizza(i)
	}
	wg.Wait()
}