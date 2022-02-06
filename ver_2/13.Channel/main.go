package main

import (
	"fmt"
	"sync"
)

func main() {
	example1()

	example2()

	// example3()
}

func example1() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Println("1 ", i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println("2 ", <-ch, "\n")
		wg.Done()
	}()

	wg.Wait()
}

func example2() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 50)

	go func(ch <-chan int) {
		// i := <-ch

		// for i := range ch {
		// 	fmt.Println("3 ", i)
		// }

		for {
			if i, isValid := <-ch; isValid {
				fmt.Println("3 ", i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		a := 43
		ch <- a
		ch <- 44
		close(ch)

		wg.Done()
	}(ch)

	wg.Wait()
	fmt.Println()
}

func example3() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Printf("Change 1: %v\n", i)
			case j := <-ch2:
				fmt.Printf("Change 2: %v\n", j)
			default:
				break
			}
		}
		wg.Done()
	}()

	go func() {
		ch1 <- 45
		close(ch1)
		ch2 <- 46
		close(ch2)
	}()
	wg.Wait()
}
