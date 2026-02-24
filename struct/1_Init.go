package main

import "fmt"

type Person struct {
	Name string
	Age int
}


func main() {
	// Full init
	p1 := Person{Name: "Alice", Age: 25}

	// Init without  assinging fields
	p2 := Person{"Bob", 30}

	// Empty struct (by default)
	p3 := Person{}

	fmt.Println(p1) // {Alice 25}
	fmt.Println(p2) // {Bob 30}
	fmt.Println(p3) //{ 0}
}
