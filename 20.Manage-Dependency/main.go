package main

// https://www.youtube.com/watch?v=-2BMNjGr2B0&list=PLVDJsRQrTUz5icsxSfKdymhghOtLNFn-k&index=1

import (
    "github.com/leekchan/accounting"
    "fmt"
)

func main() {
    ac := accounting.Accounting{Symbol: "$", Precision: 2}
    fmt.Println(ac.FormatMoney(123456789.213123))
}
