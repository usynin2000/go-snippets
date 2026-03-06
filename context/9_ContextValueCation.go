package main

import (
	"context"
	"fmt"
)

// WithValue -> IS NOT FOR optional params
// USE case: request-scoped data (auth, tracing, request metadata)

func main() {
	//  ❌ Bad: pass business data with context
	// ctx = context.WithValue(ctx, "orderID", 123)

	// ✅ Good: id for tracing
	ctx := context.WithValue(context.Background(), requestIDKey, "req-xyz")

	fmt.Println(ctx)
}
