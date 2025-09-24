// 1. Создание и добавление

// Создай пустой map[string]int.
// Добавь туда ключи "Alice" → 25, "Bob" → 30.
// Выведи всё содержимое.

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30

	fmt.Println(m)
}