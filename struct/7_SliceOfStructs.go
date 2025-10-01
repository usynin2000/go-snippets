package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
	}

	for _, p := range people {
		fmt.Println(p.Name, ":", p.Age)
	}
}