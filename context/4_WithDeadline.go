package main

import (
	"context"
	"fmt"
	"time"
)

// WithDeadline - cancelation at specifc time

func processTask(ctx context.Context, taskName string) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("Task '%s': deadline is set on %v\n", taskName, deadline.Format("21:41:02"))
	}

	for i := 1; i <= 5; i++ {
		select {
		case <- ctx.Done():
			fmt.Printf("Task '%s': shut down on step %d. The reason is %v\n", taskName, i, ctx.Err())
			return
		default:
			fmt.Printf("Task '%s': step %d\n", taskName, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	
	fmt.Printf("Task '%s': completed sucessfully\n", taskName)
}

func main() {
	// Set deadline to 2 seconds from now
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("The time now is %v\n", time.Now().Format("21:45:01"))
	fmt.Printf("Deadline: %v\n\n", deadline.Format("21:45:01"))

	processTask(ctx, "Important task")
	
	fmt.Println("\n === The programm successfuly stopped ===")
}