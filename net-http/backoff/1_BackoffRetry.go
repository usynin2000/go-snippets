package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	operation := func () error {
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

	b := backoff.NewExponentialBackOff()
	if err := backoff.Retry(operation, b); err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
