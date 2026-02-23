package main

import "fmt"


func main() {
	var ptr *int // pointer withou value -> nil

	if ptr == nil {
		fmt.Println("Poitner is empry")
	}

	x := 7

	ptr = &x

	fmt.Println("Now ptr points to:", *ptr)
}

// Pointer by default equals nil, before we assign to it address
