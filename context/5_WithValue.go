package main

import (
	"context"
	"fmt"
)

// WithValue - passing request-scoped values through context

// Define custom types for keys to avoid collisions
type contextKey string

const (
	userIDKey	contextKey = "userID"
	requestIDKey	contextKey = "requesID"
	roleKey		contextKey = "role"
)

func authenticate(ctx context.Context) context.Context {
	// add user information to context
}