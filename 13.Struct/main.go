package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student1 struct {
	id   int
	name string
	info StudentInfo
}

type Student struct {
	id   int
	name string
}

func main() {
	// named struct
	st1 := Student{
		id:   123,
		name: "Ryan",
	}

	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)
	fmt.Println()

	st2 := Student{456, "Robin"}
	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)
	fmt.Println()

	var st3 Student = struct {
		id   int
		name string
	}{
		id:   777,
		name: "Billy",
	}
	fmt.Println(st3)
	fmt.Println(st3.id)
	fmt.Println(st3.name)
	fmt.Println()

	// zero value
	var student3 Student
	student3.name = "Chaitanya"
	fmt.Println("student3 ", student3)
	fmt.Println()

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		email: "ryan@gmail.com",
		age:   27,
	}
	fmt.Println(anonymous)
	fmt.Println()

	// pointer in struct
	pointer := &Student{
		999,
		"Robin",
	}
	fmt.Println("pointer ", &pointer)
	fmt.Println((*pointer).id)
	fmt.Println((*pointer).name)
	fmt.Println(pointer.id)
	fmt.Println(pointer.name)
	fmt.Println()

	// anonymous fields
	type NoName struct {
		string
		int
	}

	n := NoName{"That string", 666}
	fmt.Println(n)
	fmt.Println()

	// long struct - Nested struct
	student2 := Student1{
		id:   123,
		name: "Brainy",
		info: StudentInfo{
			class: "ql019",
			email: "Brainy@gmail.com",
			age:   28,
		},
	}
	fmt.Println(student2)
	fmt.Println()

	st4 := Student1{
		456,
		"Robert",
		StudentInfo{
			class: "ql091",
			email: "Robert@gmail.com",
			age:   19,
		},
	}
	fmt.Println(st4)
	fmt.Println()

	st5 := Student1{
		id:   987,
		name: "Mitty",
	}
	fmt.Println("st5 ", st5)
	fmt.Println()

	// Compare 2 struct
	// 1. struct can compare when variable can checked
	type struct1 struct {
		id   int
		name string
	}
	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}

	if s1 == s2 {
		fmt.Println("s1 = s2")
	} else {
		fmt.Println("s1 != s2")
	}
	fmt.Println()

	//  2. can't be compared
	/*
		type struct2 struct {
			id   int
			name string
			info map[int]int
		}
		s3 := struct2{
			1,
			"A",
			map[int]int{
				0: 1,
			},
		}
		s4 := struct2{
			1,
			"A",
			map[int]int{
				0: 1,
			},
		}

		if s3 == s4 {
			fmt.Println("s3 = s4")
		} else {
			fmt.Println("s3 != s4")
		}
	*/

}
