package main

import "fmt"

func main() {
	// Create map with literal
	ages := map[string]int{
		"Alice": 25,
		"Bob": 30,
	}

	// Create empty map through make
	m := make(map[string]int)

	// Add element
	m["Charlie"] = 40

	fmt.Println("ages: ", ages)
	fmt.Println("m: ", m)

}
