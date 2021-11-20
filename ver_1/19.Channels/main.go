package main

import (
    "fmt"
)

// Channel
/*
    1. Unbuffered channel
    2. Buffered channel
    3. Select
    4. Close channel
*/

func main() {
    // 1. Unbuffered channel
    ch := make(chan int)

    go func() {
        ch <- 100
        fmt.Println("Sent")
    }()

    fmt.Println(<-ch)
    fmt.Println("Done", ch, "\n")

    // 2. Buffered channel
    ck := make(chan int, 2)

    ck <- 1
    ck <- 2
    // ck <- 3
    close(ck)

    fmt.Println(<-ck)
    fmt.Println(<-ck)
    fmt.Println(<-ck)
    fmt.Println(<-ck)
    fmt.Println()

    // 4. Close channel
    queue1 := make(chan int, 10)

    go func() {
        for i :=0; i<10; i++ {
            if (i > 5) {
                close(queue1)
                break
            } else {
                queue1 <- i
            }
        }
    }()

    for value := range queue1 {
        fmt.Println("close channel ", value)
    }
    fmt.Println()

    /*
      couldn't sent value to channel queue has been closed

      queue2 := make(chan int, 10)
      close(queue2)
      queue2 <- 1
    */

    // 3. Select
    queue := make(chan int)
    done := make(chan bool)

    go func() {
        for i :=0; i<10; i++ {
            queue <-i
        }

        done <- true
    }()

    for {
        select {
            case v := <-queue:
                fmt.Println(v)
            case <-done:
                fmt.Println("Done")
                return
        }
    }
    fmt.Println()
}
