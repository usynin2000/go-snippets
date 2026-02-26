package main

import "fmt"


func main() {
	a := []int{1, 2, 3}
	b := append([]int(nil), a...) // creates a copy, is considered more "go-ish" when cloning.

	b[0] = 100

	fmt.Println("a: ", a) // [1 2 3]
	fmt.Println("b: ", b) // [100, 2, 3]
}
