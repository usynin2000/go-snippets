package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	// 1. WithContext - cancel by timeout or cancelation
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	operation := func() error {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// 2. Permanent — not retry 4xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return backoff.Permanent(fmt.Errorf("client error: %d", resp.StatusCode))
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status %d", resp.StatusCode)
		}
		_, _ = io.ReadAll(resp.Body)
		return nil
	}

	b := backoff.NewExponentialBackOff(
		backoff.WithInitialInterval(100 * time.Millisecond),
		backoff.WithMaxElapsedTime(10 * time.Second),
	)
	bWithCtx := backoff.WithContext(b, ctx)

	if err := backoff.Retry(operation, bWithCtx); err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			fmt.Println("timeout")
			return
		}
		panic(err)
	}
	fmt.Println("OK")
}
