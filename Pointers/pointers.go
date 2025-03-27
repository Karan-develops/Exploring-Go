package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", num, *num)
}

func main() {
	num := 1

	fmt.Println("Memory Address:", &num)

	changeNum(&num)

	fmt.Println("In mainFunction", num)
}
