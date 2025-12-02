package main

import (
	"context"
	"fmt"
	"time"
)

// WithTimeout - automaic cancellation after timeout

func longOperation(ctx context.Context, duration time.Duration) error {
	select {
	case <-time.After(duration):
		fmt.Println("Operation completed succesfully")
		return nil
	case <- ctx.Done():
		return fmt.Errorf("Operation is cancelled: %w", ctx.Err())
	}
}

func main() {
	fmt.Println("=== Exmaple 1: Operation completes before timeout ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2 *time.Second)
	defer cancel1() // Always call cancel to release resources

	err := longOperation(ctx1, 1*time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n === Example 2: Timeout occurs ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel2()

	err = longOperation(ctx2, 3 * time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Check context errror
	if ctx2.Err() == context.DeadlineExceeded {
		fmt.Println("Context cancelled by timeout")
	}
	



}