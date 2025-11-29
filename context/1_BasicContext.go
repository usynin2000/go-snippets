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

}