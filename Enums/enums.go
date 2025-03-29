// Enumerated Types
// Go doesn't have built-in support for enums. However, you can create your own custom types to represent enums.
// Here is an example of how to create an enum in Go.
package main

import "fmt"

// iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
// It is reset to 0 when the const block starts and increments by one for each ConstSpec.
type OrderStatus int

const (
	Pending    OrderStatus = iota
	Processing             // 1
	Shipped                // 2
	Delivered              // 3
	Cancelled              // 4
	Recieved               // 5
	Confirmed              // 6
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to:", status)
}

func main() {
	// OP -> Changing Order status to: 6
	changeOrderStatus(Confirmed)
}
