package main

import (
	"fmt"
	"time"
)

// Go Does not have classes, but you can define methods on types.
// Structs are the way to define custom types in Go.

type order struct {
	id        string
	product   string
	quantity  int
	createdAt time.Time
}

// Constructor in Go

func newOrder(id string, product string, quantity int) *order {
	return &order{
		id:        id,
		product:   product,
		quantity:  quantity,
		createdAt: time.Now(),
	}
}

// Structs can have methods

// We have pass pointer to the struct to modify the struct
func (o *order) setQuantity(q int) {
	o.quantity = q
}

func (o order) getProduct() string {
	return o.product
}

func main() {
	// Structs are value types
	o := order{
		id:        "1",
		product:   "Laptop",
		quantity:  1,
		createdAt: time.Now(),
	}

	// Structs are mutable
	o.quantity = 2

	/* OP -> {1 Laptop 2 {13975070864297615848 545801 0xf539e0}} */
	fmt.Println(o)

	fmt.Println(o.getProduct()) // OP -> Laptop

	o.setQuantity(5)
	fmt.Println(o.quantity) // OP -> 5

	// Using Constructor
	o1 := newOrder("2", "Mobile", 1)
	fmt.Println(o1)         // &{2 Mobile 1 {13975072083568809512 2988701 0xaf49e0}}
	fmt.Println(o1.product) // OP -> Mobile

	// Anonymous Structs
	myOrder := struct {
		id      int
		product string
	}{
		id:      1,
		product: "Laptop",
	}

	fmt.Println(myOrder) // {1 Laptop}
}
