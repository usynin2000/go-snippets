
package main

import "fmt"

// Function types
type MathOperation func(int, int) int
type Predicate func(int) bool
type Transformer func(int) int

// We can user func as types

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// func recives func as param
func calculate(a, b int, op MathOperation) int {
	return op(a, b)
}

// Array for funcs
func applyOperations(value int, operations []Transformer) int {
	result := value
	for _, op := range operations {
		result = op(result)
	}
	return result
}

func getOperation(opType string) MathOperation {
	switch opType {
	case "add":
		return add
	case "multiply":
		return multiply
	default:
		return func(a, b int) int { return 0 }
	}
}

func mapSlice(slice []int, fn Transformer) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	// User func as value

	var op MathOperation = add
	fmt.Println("add(5, 3):", op(5, 3))

	op = multiply
	fmt.Println("multiply(5, 3)", op(5, 3))

	// pass func to another func
	result := calculate(10, 5, multiply)
	fmt.Println("calculate(10, 5, multiply):", result)

	// get func from function
	addOp := getOperation("add")
	fmt.Println("getOperation('add')(7, 3)", addOp(7, 3))

    // array of functions
	operations := []Transformer{
		func(n int) int { return n + 1},
		func(n int) int { return n * 2},
		func(n int) int { return n - 5},
	}

	final := applyOperations(10, operations)
	fmt.Println("applyOperations(10, ops):", final)

}
