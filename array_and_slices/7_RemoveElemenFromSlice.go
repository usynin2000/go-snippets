package main

import "fmt"


func main() {
	slice := []int{1, 2, 3, 4 ,5}

	index := 2 // remove element with index 2 (number 3)

	slice = append(slice[:index], slice[index + 1:]...)
	// ... takes a slice and "spreads" its elements as separate arguments of the function.

	fmt.Println(slice) // [1, 2, 4, 5]
}

// slice[:index] — all elements before the removed one.

// slice[index+1:] — all elements after the removed one.

// append(..., ......) — concatenates them into a new slice.


// In Go ... in this context is called spread operator (or "variadic expansion").
