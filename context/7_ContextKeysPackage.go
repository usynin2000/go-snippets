// File context/keys/keys.go (could be work as own package)
package keys

type contextKey string

contst (
	UserID contextKey = "userID"
	RequestID  contextKey = "requestID"
	TraceID    contextKey = "traceID"
	CorrelationID contextKey = "correlationID"
)

// In real projects keys always in  pkg/context or internal/context
