package main

import "fmt"

func main() {
	// Simple loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// Cycle with decrement
	for i := 10; i > 0; i-- {
		fmt.Printf("Reverse count: %d\n", i)
	}

	// Cycle with step
	for i := 0; i < 20; i += 2 {
		fmt.Printf("Even numbers: %d\n", i)
	}

}
