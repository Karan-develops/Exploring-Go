// Channels are used to communicate between goroutines.
// Channels are used to send and receive data between goroutines.
// Channels are type safe.
package main

import (
	"fmt"
	// "math/rand/v2"
	"time"
)

// Another Goroutine
func processNum(numChan chan int) {
	fmt.Println("Processing number:", <-numChan)
}

/*func processInfiniteNumberQueue(numInf chan int) {
	for num := range numInf {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}*/

func processSum(resultChan chan int, num1 int, num2 int) {
	numResult := num1 + num2
	resultChan <- numResult
}

// Main Fn is also a goroutine
func main() {
	// **** //
	numCh := make(chan int)

	go processNum(numCh) // Passing value to another goroutine

	numCh <- 5

	time.Sleep(time.Second * 2)

	// **** //

	// **Sending data to channel** //
	/*infCh := make(chan int)

	go processInfiniteNumberQueue(infCh)

	for {
		infCh <- rand.IntN(100)
	}*/

	// **** //

	// **Recieving data from channel** //

	resultChan := make(chan int)
	go processSum(resultChan,10,20)
	fmt.Println(<-resultChan)

	// **** //

	/*
			fatal error: all goroutines are asleep - deadlock!
			goroutine 1 [chan send]:
			main.main()
		    f:/Go/Channels/channels.go:11 +0x36
			exit status 2
	*/
	// ch := make(chan string)

	// ch <- "Karan Aggarwal"

	// msg := <-ch

	// fmt.Println(msg)
}
