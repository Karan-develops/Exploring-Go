package main

import "fmt"

// Without Generics - Repeating Code
func printIntSlices(arr []int) {
	for _, v := range arr {
		fmt.Print(v)
	}
}

func printStringSlices(arr []string) {
	for _, v := range arr {
		fmt.Print(v)
	}
}

// *********************************

// With Generics
// func printAnySlices[T any](arr []T) {
func printAnySlices[T int | string | bool](arr []T) {
	for _, v := range arr {
		fmt.Print(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"a", "b", "c", "d", "e"}
	vals := []bool{true, false, true, false, true}
	printIntSlices(nums)
	printStringSlices(names)
	printAnySlices(vals)
}
