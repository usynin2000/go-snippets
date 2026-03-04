package main

import (
	"fmt"
	"sync"
)

// All three workers work in parallel
// (or pseudo-parallel if you have one core).
// The order of output may be different each time you run it.

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", id)
}

// id - worker number (just for example).

// wg - pointer to WaitGroup, common for all goroutines.

// defer wg.Done() - guarantees that when the function exits,
//  wg.Done() will be called, decreasing the counter of active tasks by 1.

// fmt.Printf(...) - simply prints a message.

func main() {
	var wg sync.WaitGroup // var wg sync.WaitGroup - creates a waiting group.
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increases the counter of tasks by 1 (we say: "there will be another goroutine").
		go worker(i, &wg) // starts the worker function as a goroutine,
		// that is, in a separate thread of execution
	}

	wg.Wait()
	// After starting all goroutines, wg.Wait() is called - it blocks the execution of main,
	// until all goroutines call wg.Done().

	fmt.Println("all workers are done\n")

	// When the counter reaches 0, the program continues execution and prints "all workers are done".
}

// It starts three parallel goroutines
// and then waits until all of them finish their work, before printing
// all workers are done.


// Goroutine (goroutine) is a lightweight thread of execution,
// which is created with the go keyword in Go.

// 🧩 Simply put:

// When you write

// go someFunction()

// Go does not wait for someFunction() to finish -
// it starts in the background, and the main program continues working further.

// 💡 The main difference from regular threads:

// - Goroutines are very lightweight - you can start thousands (even millions) simultaneously.
// - They are managed by the Go scheduler, not the operating system.
// - One OS thread can serve many goroutines, switching between them as needed.

// (Go uses the so-called M:N scheduler, where M is the number of OS threads, and N is the number of goroutines.)
