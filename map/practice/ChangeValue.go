// 5. Замена значений

// Создай map[string]string с переводами слов:
// "cat" → "кот", "dog" → "собака".
// Замени перевод "dog" на "пёс".
package main

import "fmt"

func main() {
	m := map[string]string{
		"cat": "кот",
		"dog": "собака",
	}
	m["dog"] = "пёс"

	fmt.Println(m)
}