package main

import "fmt"

// Named return values - when we can define what values we returns
func divide(a, b float64) (quotient float64, remainder float64) {
	q := int(a / b)
	quotient = float64(q)
	remainder = a - (quotient * b)
	return // naked return (returning named values quotient and remainder)
}

// You can use only one name if only one returns
func calculate(x, y int) (result int) {
	result = x * y
	result += 10
	return // automatically return result variable
}

// Naming returned values is good for docs
func splitName(fullName string) (firstName, lastName string) {
	// You can use naked return for named values
	parts := []rune(fullName)
	mid := len(parts) / 2
	firstName = string(parts[:mid])
	lastName = string(parts[mid:])
	return 
}


func main() {
	q, r := divide(10.0, 3.0)
	fmt.Printf("10 / 3 = %.2f, remainder: %.2f\n", q, r)

	res := calculate(5, 3)
	fmt.Println("calculate(5, 3)", res)

	first, last := splitName("SergeyUsynin")
	fmt.Printf("First: %s, Last: %s\n", first, last)
}

