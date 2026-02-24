package main

import "fmt"


func main() {
	// We can make stcut without using type

	p := struct {
		Name string
		Age int
	}{
		Name: "Charlie",
		Age: 22,
	}

	fmt.Println(p)
}
