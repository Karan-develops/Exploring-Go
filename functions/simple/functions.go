package main

import "fmt"

// If type is same we can write like this
// func add(a, b int) int {
func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

// func processFunction() func(a int) int {
// 	return func(a int) int {
// 		return 10
// 	}
// }

func main() {
	fmt.Println(add(5, 3))      // 8
	fmt.Println(getLanguages()) // golang javascript c

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2) // golang javascript

	x := func(a int) int {
		return 2 + a
	}
	fmt.Println(processIt(x)) // 3

	// fmt.Println(processFunction()) // 0x67ab20
}
