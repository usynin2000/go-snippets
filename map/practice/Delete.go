// 3. Delete

// Given map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}.
// Delete "Bob".
// Print map.
package main

import "fmt"

func main() {
	m := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}

	delete(m, "Bob")

	fmt.Println(m)
}
