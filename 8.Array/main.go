package main

import "fmt"

func main() {
	// declare array - with 4 element
	var myArray [4]int

	fmt.Println(myArray)

	// add value for array
	myArray[0] = 123
	myArray[3] = 456
	fmt.Println(myArray, "\n")

	// declare array and add value
	newArray := [3]int{1, 2, 3}
	fmt.Println(newArray)
	fmt.Println("Size array ", len(newArray))

	// create array but declare size
	array2 := [...]string{"Vinfast", "Honda", "Hyundai", "Tesla"}
	fmt.Println(array2)
	fmt.Println("Size array array2 is", len(array2))
	fmt.Println(array2[3], "\n")

	// ********* ARRAY is value type but not reference type **********
	countries := [...]string{"VN", "US", "FRANCE"}
	copyCountries := countries
	fmt.Println(copyCountries)

	copyCountries[2] = "CANADA"
	fmt.Println(countries)
	fmt.Println(copyCountries, "\n")

	// loop array case 1
	for i := 0; i < len(countries); i++ {
		fmt.Println(copyCountries[i])
	}
	fmt.Println("\n")

	// loop array case 2
	for index, value := range countries {
		fmt.Printf("i= %d value= %s \n", index, value)
	}
	fmt.Println("\n")

	// blank identifier _
	for _, value := range countries {
		fmt.Printf("value= %s \n", value)
	}
	fmt.Println("\n")

	// Two dimensional arrays or 2D array - [row][column]
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{6, 5},
		{7, 8},
	}
	fmt.Println(matrix, "\n")

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " | ")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " | ")
		}
		fmt.Println()
	}
}
