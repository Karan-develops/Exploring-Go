package main

import (
	"fmt"
	"maps"
)

func main() {
	mp := make(map[string]string)
	fmt.Println(mp) // map[]

	// Setting an element
	mp["name"] = "golang"
	mp["area"] = "backend"
	mp["price"] = "Hundred"
	fmt.Println(mp) // map[area:backend name:golang price:Hundred]

	// Get element
	fmt.Println(mp["name"]) // golang
	fmt.Println(mp["area"]) // backend

	//** IMP -> If key doesn't exist in map it will return zeroed value
	// Means false for bool, 0 for int, empty for string
	fmt.Println(mp["phone"]) //

	// Length of Map
	fmt.Println(len(mp)) // 3

	// Deleting from map
	delete(mp, "price")
	fmt.Println(len(mp)) // 2

	// Clearing map -> clear(mp)

	// Other way to create map
	m := map[string]int{"price": 40, "phones": 150}
	fmt.Println(m) // map[price:40 phones:150]

	// _, ok := m["phones"]
	// return value,return bool := .....
	k, ok := m["phones"]
	fmt.Println(k) // 150
	if ok {
		fmt.Println("All Ok")
	} else {
		fmt.Println("Not Ok")
	}

	// Equality check in maps
	a := map[string]int{"phone": 100, "number": 200}
	b := map[string]int{"phone": 100, "number": 200}
	fmt.Println(maps.Equal(a, b)) // true
}
