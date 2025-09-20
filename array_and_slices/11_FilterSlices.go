package main

import "fmt"


func main() {
	names := []string{"alice", "bob", "charlie", "Dave"}

	var short []string

	for _, name := range(names) {
		if len(name) <= 4 {
			short = append(short, name)
		}
	}

	fmt.Println(short) // [bob Dave]
}