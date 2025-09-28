// 2. Индексация

// Создай строку "Hello".
// Выведи:

// первый байт,

// первый символ (через rune).
package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println("Первый бит:", s[0])
	
	runes := []rune(s)
	fmt.Println("Первый символ:", string(runes[0]))
}