// 1. Длина строки

// Создай строку "Привет, Go!".
// Выведи:

// её длину в байтах (len(s)),

// и количество символов (через []rune).

package main

import "fmt"

func main() {
	s := "Привет, Go!"
	fmt.Println("Длина в байтах", len(s))

	runes := []rune(s)
	fmt.Println("Количество символов: ", len(runes))
}