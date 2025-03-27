package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
