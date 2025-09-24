package main

import "fmt"


func main() {
	classes := map[string][]string{
		"math": {"Alice", "Bob"},
		"english": {"Charlie"},
	}

	// Добавление ученика в английский класс
	classes["english"] = append(classes["english"], "Dave")

	fmt.Println(classes)
}