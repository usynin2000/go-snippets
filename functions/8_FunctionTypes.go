package main

import "fmt"


// Function types
type MathOpertation func(int int) int
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
func calculate(a, b int, op MathOpertation) int {
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

func getOperation(opType string) MathOpertation {
	switch opType {
	case "add":
		return add
	case "multiply":
		return multiply
	default:
		return func(a, b int) int { return 0 }
	}
}