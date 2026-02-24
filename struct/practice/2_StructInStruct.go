// 🔹 Task 2. Nested structs

// Create a Person struct with fields Name and Address.
// The Address struct should contain City and Zip fields.
// Create a person "Alice" from Moscow with zip code "101000" and print their name and city.

// Expected output:

// Alice lives in Moscow

package main

import "fmt"


type Address struct {
	City string
	Zip int
}

type Person struct {
	Name string
	Address
}

func main() {
	p := Person{"Alice", Address{"Moscow", 123432}}

	fmt.Println(p.Name, "lives in", p.City)
}