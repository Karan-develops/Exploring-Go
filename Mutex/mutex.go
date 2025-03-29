// Description: This program demonstrates the use of mutex in Go.
// Mutex is used to lock the resource so that only one goroutine can access it at a time.
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	counter++
	fmt.Println("Counter:", counter)
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
}
