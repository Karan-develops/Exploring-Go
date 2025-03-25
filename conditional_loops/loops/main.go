// Loops
// In go, we have only one loop that is **("for")** loop. But we can use it in different ways to achieve the same result as other languages.
package main

import "fmt"

func main() {
	i := 1

	// 1) While loop
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// 2) Infinite Loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// 3) Classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for z := 0; z < 5; z++ {
		// break
		if z == 2 {
			continue
		}
		fmt.Println(z)
	}

	// Range (excluding 4)
	for x := range 4 {
		fmt.Println(x)
	}
}
