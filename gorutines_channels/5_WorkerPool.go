package main

import (
	"fmt"
	"time"
)


type Task struct {
	ID int
	Data string
}

// Worker is getting tasks from channel
func worker(id int, tasks <-chan Task, results chan<- string) {
	for task := range tasks {
		// Simulating task handling
		fmt.Printf("Worker %d handling task %d: %s\n", id, task.ID, task.Data)
		time.Sleep(500 * time.Millisecond)


		// Sending result
		results <- fmt.Sprintf("Worker %d completed task %d", id, task.ID)
	}
}

func main() {
	tasks := make(chan Task, 10)
	results := make(chan string, 10)

	// Making a pool with 3 workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	// Sendiing task
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- Task{ID: i, Data: fmt.Sprintf("Data of task %d", i)}
		}
		close(tasks)
	}()

	// Getting reults
	for i := 1; i <= 10; i++ {
		result := <-results
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Println("All tasks are completed")
	fmt.Println("Practical example: http handling, handling messages queue")
}