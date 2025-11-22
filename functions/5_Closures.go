package main


import "fmt"

// Closure is a function, which get variabel from outer scope
func makeCounter () func() int {
	count := 0 // func captured by closure
	return func() int {
		count++ // change captured variable
		return count
	}
}

// Closure with param
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // factor is from outer scope
	}
}