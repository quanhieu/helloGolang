// https://www.youtube.com/watch?v=XnzD1teO9fY&list=PLVDJsRQrTUz5icsxSfKdymhghOtLNFn-k&index=3
package main

import (
	"fmt"
	"runtime"
)

func g1() { fmt.Println("g1()") }
func g2() {
    fmt.Println("g2()")
//     os.openFile()
//     run login 1
//     run login 2
}
func g3() { fmt.Println("g3()") }
func g4() { fmt.Println("g4()") }

func main() {
	runtime.GOMAXPROCS(2)

	numberP := runtime.NumCPU()
	fmt.Println(numberP)

	numberP1 := runtime.GOMAXPROCS(0)
	fmt.Println(numberP1)

    go g1()
    go g2()
    go g3()
    go g4()
}

