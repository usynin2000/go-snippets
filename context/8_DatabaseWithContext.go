package main

import (
	"context"
	"fmt"
)

// Passing request-scoped data for DB layer/services


func fetchUserOrders(ctx context.Context, userID string) {
	// in logs and traces: requestID, trace_ID from ctx
	requestID := ctx.Value("requestID")
	fmt.Printf("DB query: orders for users=%s, requestID=%v\n", userID, requestID)
}

func orderService(ctx context.Context) {
	userID := ctx.Value("userID").(string)
	// context is passing for timeouts, canceling and logging
	fetchUserOrders(ctx, userID)
}


func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "req-123")
	ctx = context.WithValue(ctx, "userID", "user-456")
	orderService(ctx)
}
