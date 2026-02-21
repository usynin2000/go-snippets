package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("add: ", add(2, 3))
	q, r := divMod(10, 3)
	fmt.Println("div:", q, "mod:", r)
	
	// div: division result (quotient)
	// mod: remainder (modulo)
}
