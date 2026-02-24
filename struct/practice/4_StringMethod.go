// 🔹 The String() method

// Create a Person struct with fields Name and Age.
// Implement the String() string method for it that returns a string in the format:
// Name: Alice, Age: 25
// And print it via fmt.Println(p).
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	p := Person{"Alice", 25}
	fmt.Println(p.String())
	// This works, but Go has a special feature: if a struct implements the String() string method,
	// fmt.Println(p) calls it automatically 🎉
	fmt.Println(p)

}