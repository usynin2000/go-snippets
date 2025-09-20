package main

import "fmt"


func main() {
	a := []int{4, 5}
	b := []int{1, 2, 3}

	b = append(b, a...) // эквивалентно: append(b, 4, 5)

	fmt.Println(b)
}