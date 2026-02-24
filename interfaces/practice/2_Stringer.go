// 🔹 Task 2. Stringer

// Create a Person struct {Name string; Age int}.
// Make it so that when printing via fmt.Println(p) it outputs:

// Alice (25 years old)
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
	p := Person{"Alice", 25}

	fmt.Println(p)
}