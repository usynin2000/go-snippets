package main

import "fmt"

// Interface constraint
type Printer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %s", p.Name)
}

// Function gets as param only types which implements by Printer
func PrintAll[T Printer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}


func main() {
	people := []Person{
		{Name: "Alice"},
		{Name: "Bob"},
	}
	
	PrintAll(people)
}