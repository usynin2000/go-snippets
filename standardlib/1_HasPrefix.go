package main

import "fmt"

// HasPrefix tests whether the string s begins with prefix

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func main() {

	hasPrefix := HasPrefix("Serega", "Ser")
	fmt.Println(hasPrefix) // true 


	hasPrefix = HasPrefix("ololo", "he")
	fmt.Println(hasPrefix) // false

}