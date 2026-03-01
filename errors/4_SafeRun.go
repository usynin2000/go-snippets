package main

import "fmt"

func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered after:", r)
		}
	}()

	fmt.Println("We are executing...")
	panic("fatal error")
}

func main() {
	safeRun()
	fmt.Println("The program continues to work")
}

// We are executing...
// Recovered after: fatal error
// The program continues to work
