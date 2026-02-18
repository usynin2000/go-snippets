package main

import "fmt"

func main() {
	// Slice cycle
	fruits := []string{"яблоко", "банан", "апельсин"}

	for index, fruit := range fruits {
		fmt.Printf("Индекс %d: %s\n", index, fruit)
	}

	// Cycel for values only
	for _, fruit := range fruits {
		fmt.Printf("Фрукт: %s\n", fruit)
	}

	// Cycle for indexes only ( no need to add _ after index)
	for index := range fruits {
		fmt.Printf("Индекс: %d\n", index)
	}

	// Cycle on map
	colors := map[string]string{
		"red": "красный",
		"blue": "синий",
		"green": "зеленый",
	}

	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}

	// Cycle on string
	text := "Hello"
	for i, char := range text {
		fmt.Printf("Position %d: %c (код: %d)\n", i, char, char)
	}


}
