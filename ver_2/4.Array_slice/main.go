package main

import (
	"fmt"
)

func main() {
	// Array
	// points := [3]int{7, 10, 5}
	// points := [...]int{7, 10, 5, 6}
	var points [3]int
	points[0] = 10
	points[1] = 2
	points[2] = 6

	fmt.Printf("%v , %T\n", points, points)
	fmt.Printf("%v , %T\n", points[2], points)
	fmt.Println(len(points))

	// copy array
	a := [3]int{2, 5, 10}
	b := a  // copy value at new address
	c := &a // copy value but point to a address
	b[0] = 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Slice
	d := []int{1, 3, 7, 9, 11, 13, 15, 17, 19, 21}
	e := d
	e[0] = 20
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(len(d)) // length
	fmt.Println(cap(d)) // storage of array

	f := d[:]
	g := d[3:]
	h := d[:6]
	i := d[3:6]
	fmt.Printf("\nd %v , %v, %v\n", d, len(d), cap(d))
	fmt.Printf("f %v , %v, %v\n", f, len(f), cap(f))
	fmt.Printf("g %v , %v, %v\n", g, len(g), cap(g))
	fmt.Printf("h %v , %v, %v\n", h, len(h), cap(h))
	fmt.Printf("i %v , %v, %v\n", i, len(i), cap(i))

	i[1] = 100
	fmt.Printf("\nd %v , %v, %v\n", d, len(d), cap(d))
	fmt.Printf("f %v , %v, %v\n", f, len(f), cap(f))
	fmt.Printf("g %v , %v, %v\n", g, len(g), cap(g))
	fmt.Printf("h %v , %v, %v\n", h, len(h), cap(h))
	fmt.Printf("i %v , %v, %v\n", i, len(i), cap(i))

	j := make([]int, 10, 20)
	j = append(j, 1)
	fmt.Printf("\nj %v , %v, %v\n", j, len(j), cap(j))

	k := make([]int, 0)
	k = append(j, 1)
	k = append(j, 2)
	k = append(j, 3)
	k = append(j, []int{11, 12, 13, 14, 15, 16}...)
	fmt.Printf("k %v , %v, %v\n\n", k, len(k), cap(k))

	l := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("l ", l)
	m := append(l[:2], l[4:]...)
	fmt.Println("l ", l)
	fmt.Println("m ", m)

}

/*

 */
