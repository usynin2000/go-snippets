package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println("a && b", a && b) // false
	fmt.Println("a || b", a || b) // true
	fmt.Println("!a", !a) // false
}

// && = И
// || = ИЛИ
