package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4, 5}

	a = append(a, b...)

	fmt.Println(a) // [1, 2, 3, 4, 5]
}