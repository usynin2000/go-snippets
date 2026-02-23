package main

import "fmt"

func createPointer(val int) *int {
	return &val
}

func main() {
	ptr := createPointer(99)
	fmt.Println("Value", *ptr)
}
