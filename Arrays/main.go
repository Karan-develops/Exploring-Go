package main

import "fmt"

func main() {
	var nums [3]int

	nums[1] = 10

	// Getting Array's Length
	fmt.Println(len(nums))
	fmt.Println(nums)

	// Declaring and assigning values to arra y
	nums2 := [4]int{1, 2, 3, 4}
	fmt.Println(nums2)

	// 2D - Arrays
	nums3 := [2][2]int{{0, 1}, {1, 0}}
	fmt.Println(nums3)
}
