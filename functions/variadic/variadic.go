package main

import "fmt"

// Similar to var-args
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(1, 2, 3, 50, "Hello")
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	arr := []int{1, 2, 3}
	fmt.Println(sum(arr...)) // 6
}
