// Des: This program demonstrates the use of goroutines in Go.
//
//	Goroutines are lightweight threads of execution that run concurrently.
//	They are used to achieve parallelism in Go.
//	In this program, we have a task function that prints the task id.
//
// Just use the go keyword followed by the function call to run the function as a goroutine.
package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing Task:", id)
}

func main() {
	for i := range 10 {
		go task(i)
	}
	time.Sleep(time.Second * 2)
}
