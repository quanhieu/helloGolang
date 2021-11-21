package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	number   int
	name     string
	isMale   bool
	subjects []string
}

type People struct {
	name string
	age  int
}
type StudentEmbedded struct {
	People
	number   int `Maximum can't over 100`
	subjects []string
}

func main() {
	// Map
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int{
		"Hieu": 20,
		"Tom":  25,
		"Mike": 29,
	}
	studentNameAgeMap["David"] = 24
	studentNameAgeMap["Hieu"] = 21

	fmt.Println("studentNameAgeMap ", studentNameAgeMap, len(studentNameAgeMap))
	fmt.Println(studentNameAgeMap["Hieu"])

	delete(studentNameAgeMap, "Tom")
	fmt.Println("deleted ", studentNameAgeMap)

	// ageTom, isExist := studentNameAgeMap["Tom"]
	_, isExist := studentNameAgeMap["Tom"]
	fmt.Println("ageTom ", isExist)

	copyMap := studentNameAgeMap
	fmt.Println(studentNameAgeMap)
	fmt.Println(copyMap)

	copyMap["Mike"] = 21
	fmt.Println(studentNameAgeMap)
	fmt.Println(copyMap)

	// Struct
	studentHieu := Student{
		number: 3,
		name:   "Hieu",
		isMale: true,
		subjects: []string{
			"Math", "English", "Computer",
		},
	}
	fmt.Println("\nstudentHieu ", studentHieu)
	fmt.Println("studentHieu ", studentHieu.name)

	studentEmpty := Student{}
	studentEmpty.number = 19
	studentEmpty.name = "Teo"
	studentEmpty.isMale = false
	studentEmpty.subjects = []string{
		"Math", "English", "Computer",
	}
	fmt.Println("\nstudentEmpty ", studentEmpty)

	studentStruct := struct{ name string }{
		name: "Struct_name",
	}
	copyStudent := studentStruct
	fmt.Println("\nstudentStruct ", studentStruct)
	fmt.Println("copyStudent ", copyStudent)

	copyStudent.name = "Copy"
	fmt.Println("\nstudentStruct ", studentStruct)
	fmt.Println("copyStudent ", copyStudent)

	copyStudentTwo := &studentStruct
	copyStudentTwo.name = "Copy_Two"
	fmt.Println("\nstudentStruct ", studentStruct)
	fmt.Println("copyStudentTwo ", copyStudentTwo)

	// Embeded
	studentEmbedded := StudentEmbedded{}
	studentEmbedded.age = 19
	studentEmbedded.name = "Teo"
	studentEmbedded.number = 123
	studentEmbedded.subjects = []string{
		"Math", "English", "Computer",
	}
	fmt.Println("\nstudentEmbedded ", studentEmbedded)

	// Tag
	t := reflect.TypeOf(studentEmbedded)
	field, _ := t.FieldByName("number")
	fmt.Println("\n field ", field)
}

/*

 */
