package main

import "fmt"

type Car struct {
	Brand string
	Year int
}

func main() {
	c := Car{Brand: "Toyota", Year: 2020}

	// Чтение
	fmt.Println(c.Brand)

	// Изменение 
	c.Year = 2023
	fmt.Println(c.Year)
}