package main

import (
	"context"
	"fmt"
)

// Basic context creation and usage

func main() {
	// Background context - root context, never cancelled

	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO context - use when unsure which context to use
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	// Check if context is done (it's not)
	select {
	case <- ctx.Done():
		fmt.Println("Context is done")
	default:
		fmt.Println("Context is still active")
	}

	// Context values
	fmt.Printf("Context error: %v\n", ctx.Err())
	
	deadline, _ := ctx.Deadline()
	fmt.Printf("Context deadline: %v\n", deadline)



}