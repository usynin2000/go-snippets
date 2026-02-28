package main

import "fmt"


func main() {
	s := "молоко"
	runes := []rune(s)
	runes[0] = 'т' // Single quotes '...' is a rune (rune, i.e. int32, code of character in Unicode). For example:
	// This works because runes is []rune, and each element is a rune.
	// So exactly what is expected is an integer value representing a character.
	// If you wrote: runes[0] = "т"  // string
	// Go would compile an error because "т" is a string (string), not a rune.
	// You cannot directly put a string into an []rune element.
	s = string(runes)
	fmt.Println("s:", s) // толоко
}

// Strings in Go are immutable, so that's why.
