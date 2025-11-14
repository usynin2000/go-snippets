package main

import "fmt"

// T (type parameter)
// comparable - constraint, it means that T has to be comparable

func FindInSlice[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	// Works with int
	ints := []int{10, 20, 30, 40, 50}
	fmt.Println(FindInSlice(ints, 30))

	// Works with strings
	strings := []string{"apple", "banana", "cherry"}
	fmt.Println(FindInSlice(strings, "banana"))

	// Works with float64
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println(FindInSlice(floats, 2.2))

	// Doesn't work with slices
	// slices  := [][]int{{1, 2}, {3, 4}}
	// fmt.Println(FindInSlice(slices, []int{1, 2})) // []int does not satisfy comparablecompilerInvalidTypeArg
}