package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// Make var of type Person
	p := Person{Name: "Serega", Age: 24}

	// Make pointer to var structure
	ptr := &p

	// Change field by using pointer
	ptr.Age = 150

	// No need to use pointer dereference like.
	//  (*ptr).Age == ptr.Age

	fmt.Println("Name:", p.Name)
	fmt.Println("Age", p.Age)

}
