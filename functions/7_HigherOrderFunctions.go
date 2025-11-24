package main

import "fmt"

// Higher order function - fuctions that recieves or returns another functions

// Map - make func to every slice element
func mapInts(slice []int, fn func(int) int) (result []int) {
	result = make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return 
}

// Filter - filtering element of slice by condition
func filterInts(slice []int, fn func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reducde - reduse slice to one value 
func reduceInts(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func makeValidator(min, max int) func(int) bool {
	return func(n int) bool {
		return n >= min && n <= max
	}
}

// func that recives many funcs
func processWithPipeline(value int, functions ...func(int) int) int {
	result := value
	for _, fn := range functions {
		result = fn(result)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Map: double every number
	doubled := mapInts(
		numbers, func(n int) int {
			return n * 2
		},
	)
	fmt.Println("Doubled", doubled)

	// Filter: leaves only evens
	evens := filterInts(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("Evens:", evens)

	// Reduce: to summarize all numbers
	sum := reduceInts(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("SUM:", sum)

	// Validator

	isValidAge := makeValidator(18, 65)
	fmt.Println("Age 25 valid:", isValidAge(25))
	fmt.Println("Age 70 valid:", isValidAge(70))

	// Functions Pipeline
	addOne := func(n int) int {return n + 1}
	multipleByTwo := func(n int) int { return n * 2}
	square := func(n int) int { return n * n}

	result := processWithPipeline(3, addOne, multipleByTwo, square)
	fmt.Println("Pipeline(3):", result)



}