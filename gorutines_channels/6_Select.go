package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Example 1: select for reading from several chanlles ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// select is waiting for channel to be ready
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Got from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Got from ch2: %s\n", msg2)
		}
	}

	fmt.Println("\n=== Example 2: select with timeout ===")
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Data recieved"
	}()

	select {
	case msg := <- ch:
		fmt.Printf("Success: %s\n", msg)
	
	// If 2 second pass, then we need to use this case
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Channel did not respond in 2 second")
	}

	fmt.Println("\n=== Example 3: select with default (not blocked) ===")
	nonBlocking := make(chan string)

	select {
	case msg := <-nonBlocking:
		fmt.Printf("Got %s\n", msg)
	default:
		fmt.Println("Channel is empty, however programm is not blocked")
	}

	// sending data
	go func() {
		nonBlocking <- "Here is data"
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case msg := <-nonBlocking:
		fmt.Printf("Now we got it: %s\n", msg)
	default:
		fmt.Println("Channel is still empty")
	}
}

// What it shows:
// 	select to work with several channels
// 	timeouts using time.After
// 	non-blocking opertation by using default