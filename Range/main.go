package main

import "fmt"

// Iterating over data structures
func main() {

	// 1) Slices
	nums := []int{5, 6, 7}

	/* OP
	0 -> 5
	1 -> 6
	2 -> 7*/
	for i, num := range nums {
		fmt.Println(i, "->", num)
	}

	// 2) Maps
	mp := map[string]int{"price": 100, "number": 150}

	/* OP
	price -> 100
	number -> 150
	*/
	for key, val := range mp {
		fmt.Println(key, "->", val)
	}

	// 3) Strings
	/* OP
	0 -> 103
	1 -> 111
	2 -> 108
	3 -> 97
	4 -> 110
	5 -> 103
	*/
	for i, ch := range "golang" {
		fmt.Println(i, "->", ch)
	}
	/* OP
	0 -> g
	1 -> o
	2 -> l
	3 -> a
	4 -> n
	5 -> g
	*/
	for i, ch := range "golang" {
		fmt.Println(i, "->", string(ch))
	}
}
