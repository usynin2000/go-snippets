package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)


/// HTTP Middleware pattern with con text like Echo, Ghin, CHi


type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userKey contextKey = "user"
)


func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := fmt.Sprintf("%d", time.Now().UnixNano())
		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext((ctx)))
	})
}


func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userID := r.Header.Get("X-User-ID")
			if userID == "" {
				userID = "anonymous"
			}
			ctx := context.WithValue(r.Context(), userKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID := ctx.Value(requestIDKey).(string)
	user := ctx.Value(userKey).(string)
	// Why do we need to user .(string)
	// Because context.Value(key) returns interface{} and we need to be sure
	// That our value will be string
	fmt.Fprintf(w, "RequestID: %s, User: %s\n", reqID, user)
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	handler := requestIDMiddleware((authMiddleware(mux)))

	fmt.Println("Server Starting on :8000")
	http.ListenAndServe(":8080", handler)
}
