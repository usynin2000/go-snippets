package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)) // making new slice with the same length

	copy(b, a)  // copying data
	b[0] = 100

	fmt.Println("a: ", a) // [1, 2, 3]
	fmt.Println("b: ", b) // [100, 2, 3]
}
// 💡 copy creates a new slice with separate memory.
