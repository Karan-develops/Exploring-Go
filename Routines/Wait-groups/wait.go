package main

import (
	"fmt"
	"sync"
)

// defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.

// Pass the WaitGroup as a pointer to the task function.
func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing Task:", id)
}

func main() {
	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
