package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}

	// Update value
	ages["Alice"] = 26
	// Delete element
	delete(ages, "Bob")
	fmt.Println(ages)
}
