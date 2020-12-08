package main

/*
	- Concurrency is about "dealing" with lots of things at once - 'Rob Pike'
	- Parallelism is about "doing" lots of things at once (with condition CPU need to upper 2 core) - 'Rob Pike'
	- Do not communicate by sharing memory; instead, share memory by communicating
	https://www.youtube.com/watch?v=KmJ-Phn49jA&list=PLVDJsRQrTUz5icsxSfKdymhghOtLNFn-k&index=5
	https://www.youtube.com/playlist?list=PLlahAO-uyDzIVzBvRKwKUDjj2Iaq-5W9l
*/

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
