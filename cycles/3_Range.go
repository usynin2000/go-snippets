package main

import "fmt"

func main() {
	// Цикл по слайсу
	fruits := []string{"яблоко", "банан", "апельсин"}

	for index, fruit := range fruits {
		fmt.Printf("Индекс %d: %s\n", index, fruit)
	}

	// Цикл только по значиням
	for _, fruit := range fruits {
		fmt.Printf("Фрукт: %s\n", fruit)
	}

	// Цикл только по индексам
	for index := range fruits {
		fmt.Printf("Индекс: %d\n", index)
	}

	// Цикл по мапе
	colors := map[string]string{
		"red": "красный",
		"blue": "синий",
		"green": "зеленый",
	}

	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}

	// Цикл по строке
	text := "Привет"
	for i, char := range text {
		fmt.Printf("Позиция %d: %c (код: %d)\n", i, char, char)
	}


}