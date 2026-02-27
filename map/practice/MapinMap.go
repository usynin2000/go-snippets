// 6. Nested map

// Create map[string]map[string]int, where key is name of person, and value is map "subject → score".
// For example:

// {
//   "Alice": {"math": 5, "english": 4},
//   "Bob": {"math": 3, "english": 5}
// }


// Выведи оценку Боба по английскому.
package main

import "fmt"

func main() {
	m := map[string]map[string]int{
		"Alice": {"math": 5, "english": 4},
		"Bob": {"math": 3, "english": 5},
	}

	fmt.Println("Bob's score in english", m["Bob"]["english"])
}
