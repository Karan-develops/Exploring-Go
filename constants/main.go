package main

import "fmt"

const age = 20

/*
We can't use type inference := outside functions (global scope)
age:=20 -> error
*/

func main() {
	const name = "ABC"

	// We can't use := with constants
	// const age:=15

	// Constant Group
	const (
		port   = 3000
		host   = "localhost"
		server = "http"
	)

	fmt.Println(port, host, server)
}
