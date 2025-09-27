package main

import "fmt"


func main() {
	s := "молоко"
	runes := []rune(s)
	runes[0] = 'т'
	s = string(runes)
	fmt.Println("s:", s) // толоко
}

// Строки в Go иммутабельные, поэтому так.