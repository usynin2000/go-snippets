package main

import (
	"context"
	"fmt"
	"time"
)

// WithCancel - manual canellation of operations

func worker(ctx context.Context, id int) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Worker %d: got canellation signal: %v\n", id, ctx.Err())
			return
	default:
		fmt.Printf("Worker %d: is working ...\n", id)
		time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Starting 3 workers
	for i:= 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// Let them work for 2 seconds
	time.Sleep(2 * time.Second)

	// Cancel all workers
	fmt.Println("\n=== Cancel all operations ===")
	cancel()

	// Give time to see cancellation messages
	time.Sleep(1 * time.Second)
	fmt.Println("\n === The application is completed. ===")
}