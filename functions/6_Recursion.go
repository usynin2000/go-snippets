package main

import "fmt"

// Recursion is when func calling itself
func factorial(n int) int {
	if n <= 1 {
		return 1 // Basic event
	}
	return n * factorial(n-1) // recursive call
}

// Recursion fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n // basic case: fib(0)=0, fib(1)=1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Factorial(5): ",  factorial(5)) // 120

	fmt.Println("Fibonacci(0-9): ",)
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}