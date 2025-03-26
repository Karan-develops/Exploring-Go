package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 5
	// Go doesn't require break statement
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("LOL")
	}

	// Multiple Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// type switch
	// OLD
	// whoAmI := func(i interface{}) {
	// Modern
	// whoAmI := func(i any) {

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("This is an integer")
		case string:
			fmt.Println("This is a String")
		case bool:
			fmt.Println("This is Boolean")
		default:
			fmt.Println("IDK", t)
		}
	}
	whoAmI(0.1)
}
