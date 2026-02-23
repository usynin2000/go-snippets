package main

import "fmt"

func main() {
	x := 5

	p1 := &x // Pointer to x

	p2 := &p1 // Pointer to pointer

	fmt.Println("x: ", x)
	fmt.Println("*p1: ", *p1)
	fmt.Println("**p2: ", **p2)

	**p2 = 42
	fmt.Println("New value of x: ", x)
}
