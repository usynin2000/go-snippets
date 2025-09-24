package main

import "fmt"

func main() {
	var mNil map[string]int

	fmt.Println("Is nil?", mNil == nil) // true

	// mNil["x"] = 10 // panic: assignment to entry in nil map

	// Нужно инициализировать через make
	mok := make(map[string]int)
	mok["x"] = 10
	fmt.Println(mok)
}