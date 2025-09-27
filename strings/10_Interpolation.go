package main

import "fmt"


func main() {
	name := "Alice"
	age := 30

	s := fmt.Sprintf("My name is %s and I am %d years old.", name, age)

	fmt.Println(s)
}

// Через fmt.Sprintf можно вставлять значения внутрь строки