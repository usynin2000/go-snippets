package main

import "fmt"


func main() {
	classes := map[string][]string{
		"math": {"Alice", "Bob"},
		"english": {"Charlie"},
	}

	// Add student to english class
	classes["english"] = append(classes["english"], "Dave")

	fmt.Println(classes)
}
