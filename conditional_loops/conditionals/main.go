// Go doesn't have ternary operator, so we can't use it in Go.
// But we can use if-else construct to achieve the same result.
package main

import "fmt"

func main() {

	age := 21
	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else if age == 18 {
		fmt.Println("You are Just eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote!")
	}

	var role = "admin"
	var hasPermission bool = true

	if role == "admin" && hasPermission {
		fmt.Println("You can access the admin panel!")
	} else {
		fmt.Println("You can't access the admin panel!")
	}

	// We can declare a variable inside if construct also
	if x := 10; x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}
}
