package main

import (
	"context"
	"fmt"
)

// WithValue - passing request-scoped values through context

// Define custom types for keys to avoid collisions
type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requesID"
	roleKey      contextKey = "role"
)

func authenticate(ctx context.Context) context.Context {
	return context.WithValue(ctx, userIDKey, "user-123")
}

func setRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestIDKey, "req-abs-def")
}

func setRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}


// Helper functions - is good to extract values
func getUserID(ctx context.Context) (string, bool) {
	v  := ctx.Value(userIDKey)
	if v == nil {
		return "", false
	}

	id, ok := v.(string)
	return id, ok
}

func getRequestID(ctx context.Context) string {
	v := ctx.Value(requestIDKey)
	if v == nil {
		return ""
	}
	return v.(string)
}

func handler(ctx context.Context) {
	userID, ok := getUserID(ctx)
	if !ok {
		fmt.Println("User not authenticated")
		return
	}

	fmt.Printf(
		"Processing request for user: %s, requestID: %s\nwith role %s",
		userID, getRequestID(ctx), ctx.Value(roleKey),
	)
}


func main() {
	// emulate middleware

	ctx := context.Background()
	ctx = setRequestID(ctx)
	ctx = authenticate(ctx)
	ctx = setRole(ctx, "admin")

	handler(ctx)
}



