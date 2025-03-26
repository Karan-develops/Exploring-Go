package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums2 []int
	fmt.Println(nums2 == nil) // true
	fmt.Println(len(nums2))   // 0

	// Creating a slice
	/*
		// var nums = make([]int, 2, 5)
		This will create a slice with two starting 2 zeros of cap 5 [0,0,....]
	*/

	// This slice will not have 0s in starting
	var nums3 = make([]int, 0, 5)
	fmt.Println(nums3) // [0,0]
	// Capacity -> maximum number of elements can fit
	fmt.Println(cap(nums3)) // 5

	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)
	nums3 = append(nums3, 3)
	nums3 = append(nums3, 4)

	fmt.Println(cap(nums3)) // 10 (automatically increased capacity)

	// One more way to create a slice
	nums4 := []int{}
	fmt.Println(cap(nums4)) // 0

	// Copy functions
	var arr = make([]int, 0, 5)
	arr = append(arr, 2)
	arr = append(arr, 3)
	var arr2 = make([]int, len(arr))

	copy(arr2, arr)
	fmt.Println(arr, arr2) // [2 3] [2 3]

	// Slice operator
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[0:2]) // [1 2]
	fmt.Println(a[1:])  // [2 3 4 5]
	fmt.Println(a[:3])  // [1 2 3]

	// Slice operations
	b := []int{1, 2, 3}
	c := []int{1, 2, 3}

	fmt.Println(slices.Equal(b, c)) // true

	// 2D Slices
	lol := [][]int{{1, 2}, {2, 3}}
	fmt.Println(lol) // [[1 2] [2 3]]
}
