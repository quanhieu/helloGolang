package main

import (
	"fmt"
	"sync"
	"time"
)

var newWg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func main() {
	go count("Sheep")
	count("fish")
	time.Sleep(time.Second * 2)
	fmt.Println()

	//
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		count("Sheep")
		wg.Done()
	}()

	go func() {
		count("Fish")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done\n")

	// RWMutex
	for i := 0; i < 10; i++ {
		newWg.Add(2)

		// lock Read
		m.RLock()
		go sayHello()

		// lock write
		m.Lock()
		go increment()
	}
	newWg.Wait()
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

//
func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	newWg.Done()
}

func increment() {
	counter++
	m.Unlock()
	newWg.Done()
}
