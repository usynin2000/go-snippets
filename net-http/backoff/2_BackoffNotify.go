package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	operation := func() error {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status %d", resp.StatusCode)
		}
		_, _ = io.ReadAll(resp.Body)
		return nil
	}

	notify := func(err error, d time.Duration) {
		log.Printf("error: %v, retry in %v", err, d)
	}

	b := backoff.NewExponentialBackOff(
		backoff.WithInitialInterval(200 * time.Millisecond),
		backoff.WithMaxElapsedTime(5 * time.Second),
	)
	if err := backoff.RetryNotify(operation, b, notify); err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
