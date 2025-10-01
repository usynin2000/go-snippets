package main

import "fmt"

type Address struct {
	City string
	Index int
}

type User struct {
	Name string
	Address Address
}

func main() {
	u := User{
		Name: "Alice",
		Address: Address{
			City: "Moscow",
			Index: 12345,
		},
	}

	fmt.Println(u.Name,  "lives in", u.Address.City)
}