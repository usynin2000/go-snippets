// 2. Read and check

// Given map[string]int{"Alice": 25, "Bob": 30}.
// Try to get age "Charlie".
// First, simply print the result.
// Then, do it through val, ok := m["Charlie"].

package main

import "fmt"

func main() {
	m := map[string]int{"Alice": 25, "Bob": 30}

	fmt.Println(m["Charlie"])

	val, ok := m["Charlie"]
	if ok {
		fmt.Println("Charlise есть, вот возраст его", val)
	} else {
		fmt.Println("We don't know who Charlie is and how old he is.")
	}

}
