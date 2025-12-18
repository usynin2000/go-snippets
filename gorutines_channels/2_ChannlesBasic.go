package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("== Example 1: Simple data transfer ===")

	// Make a chanel for data transfer
	ch := make(chan string)

	// Goroutine sends data to channel
	go func() {
		ch <- "Hello from goroutine!"
		fmt.Println("Data sent")
	}()
	
	// Reading datat from channel block execution until something is sent to it
	message := <-ch
	fmt.Printf("Got from channel: %s\n\n", message)

	fmt.Println("=== Example 2: Sending numbers ===")
	numbers := make(chan int)

	// Sending numbers
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("Sent: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(numbers) // Closing channel after sending all data
	}()

	fmt.Println("Reading all numbers:")
	
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers type: %T\n", numbers) // chan int

	for num := range numbers {
		fmt.Printf("Got from channel: %d\n", num)
	}
	fmt.Println("All numbers received\n")

	fmt.Println("=== Exmaple 3: Two-way communication ===")
	request := make(chan string)
	response := make(chan string)

	// Goroutine-handler
	go func () {
		for req := range request {
			response <- fmt.Sprintf("Response to: %s", req)
		}
	}()

	request <- "Request 1"
	fmt.Printf("Response: %s\n", <-response)

	request <- "Request 2"
	fmt.Printf("Response: %s\n", <-response)

	close(request)

}
// Why we need channels? Can we just store the result of the goroutine in a variable?

// Goroutines run in parallel, and if you just try to write the result of the goroutine into one variable (for example, a global variable), you will get two main problems:

// 1. No synchronization of access:
// If several goroutines write/read from one variable at the same time, data races occur. This leads to errors that are very difficult to track.

// 2. How to know when the result is ready?
// Suppose one goroutine processes data, and the main function waits for the result. 
// You need to somehow "have a signal" when the work is done — and that's where channels come in perfectly. 
// Reading from a channel blocks execution until something is sent to it — and vice versa.


// > (Do not communicate by sharing memory; instead, share memory by communicating)