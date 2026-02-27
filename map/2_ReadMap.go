package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25}

	// Access by key
	fmt.Println("Alice: ", ages["Alice"])

	// If key is not found, return "zero value" (for int it is 0)
	fmt.Println("Bob (no key): ", ages["Bob"])

	// check if key exists
	val, ok := ages["Bob"]
	if ok {
		fmt.Println("Bob found:", val)
	} else {
		fmt.Println("Bob not found")
	}
}
