package main

import "fmt"

// Filter function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)

		}
	}
	return result
}

func main() {
	// Filtering numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evens)

	words := []string{"apple", "a", "banana", "cat", "dog"}
	longWords := Filter(words, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println("Long worlds:", longWords)

}
