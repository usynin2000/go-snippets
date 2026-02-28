package main

import "fmt"


func main() {
	s := "hello"
	fmt.Println(s[:2]) // he (only with english characters)

	ru :=[]rune("Привет")
	fmt.Println(string(ru[:3])) // При (for russian since there are more bytes, you need to use runes)
}

// Technically you can take slices like an array, but you need to be careful with multi-byte characters!
