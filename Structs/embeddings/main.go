package main

import "fmt"

type customer struct {
	name string
	age  int
}

// It is like inherting customer struct
type order struct {
	id       string
	product  string
	quantity int
	customer
}

func main() {

	newOrder := order{
		id:       "1",
		product:  "Laptop",
		quantity: 1,
	}
	// {1 Laptop 1 { 0}}
	// { 0}
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)

	// OP
	// {karan 20}
	// {1 Laptop 1 {karan 20}}
	fmt.Println(customer{"karan", 20})
	fmt.Println(order{customer: customer{"karan", 20}, id: "1", product: "Laptop", quantity: 1})
}
