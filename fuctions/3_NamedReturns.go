package main

import "fmt"

// Named return values - when we can define what values we returns
func divide(a, b float64) (quotient float64, remainder float64) {
	q := int(a / b)
	quotient = float64(q)
	remainder = a - (quotient * b)
	return // naked return (returning named values quotient and remainder)
}


func main() {
	q, r := divide(10.0, 3.0)
	fmt.Printf("10 / 3 = %.2f, remainder: %.2f\n", q, r)
}

