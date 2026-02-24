package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Employee struct {
	Person // embedding
	Position string
}

func main() {
	 e := Employee {
		Person: Person{Name: "Alice", Age: 25},
		Position: "Developer",
	 }

	 fmt.Println(e.Name) // we can access to field Name without need to specify that it is another struct
	 fmt.Println(e.Position)  // own field
}
