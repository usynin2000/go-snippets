// 1. Create and add

// Create empty map[string]int.
// Add keys "Alice" → 25, "Bob" → 30.
// Print all content.

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30

	fmt.Println(m)
}
