package main

import "fmt"

func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	Describe(10)
	Describe("Go")
	Describe(3.14)
}