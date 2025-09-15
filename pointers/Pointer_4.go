package main

import "fmt"

func main() {
	x := 5

	p1 := &x // указатель на x

	p2 := &p1 // указатель на указатель

	fmt.Println("x: ", x)
	fmt.Println("*p1: ", *p1)
	fmt.Println("**p2: ", **p2)

	**p2 = 42
	fmt.Println("Новое значение x: ", x)
}

// Когда это полезно
