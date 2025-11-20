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

// Worker 3 handling task 3: Data of task 3
// Worker 2 handling task 1: Data of task 1
// Worker 1 handling task 2: Data of task 2
// Worker 1 handling task 4: Data of task 4
// Worker 3 handling task 5: Data of task 5
// Result: Worker 1 completed task 2
// Result: Worker 3 completed task 3
// Result: Worker 2 completed task 1
// Worker 2 handling task 6: Data of task 6
// Worker 2 handling task 7: Data of task 7
// Worker 1 handling task 8: Data of task 8
// Worker 3 handling task 9: Data of task 9
// Result: Worker 2 completed task 6
// Result: Worker 1 completed task 4
// Result: Worker 3 completed task 5
// Worker 1 handling task 10: Data of task 10
// Result: Worker 1 completed task 8
// Result: Worker 3 completed task 9
// Result: Worker 2 completed task 7
// Result: Worker 1 completed task 10
// All tasks are completed
// Practical example: http handling, handling messages queue

// Controlled concurrency with worker pool