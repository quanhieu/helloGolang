package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	// zero value
	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("myMap != nil")
	}
	fmt.Println()

	var myMap1 map[string]int
	fmt.Println(myMap1)

	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}
	fmt.Println()

	// declare variable
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	// add element in map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	// delete element in map = delete(map, key)
	delete(myMap2, "key2")
	fmt.Println(myMap2)

	// length
	fmt.Println("length ", len(myMap2))

	// Map is reference type
	myMap3 := myMap2
	fmt.Println(myMap3)

	myMap3["key5"] = 1000
	fmt.Println("myMap2 ", myMap2)
	fmt.Println("myMap3 ", myMap3)

	// access element in map
	value := myMap2["key5"]
	fmt.Println(value)

	// check exist key in map
	value, found := myMap2["key1000"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Cannot find value of key1000")
	}

	// cannot compare == in map
	/*
		if myMap2 == myMap3 {
			fmt.Println("Compare map")
		}
	*/
}
