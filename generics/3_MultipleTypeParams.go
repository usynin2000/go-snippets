package main

import "fmt"

func Swap[T any, U any](a T, b U) (U, T) {
	return b, a
}

func main() {
	a, b := Swap("Hello", 42)
	fmt.Println(a, b) // 42 Hello
}