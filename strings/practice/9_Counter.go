// ⭐ 9. Подсчёт символов

// Дана строка "банан".
// Посчитай, сколько раз встречается каждая буква, сохрани результат в map[rune]int.
package main

import (
	"fmt"
)

func main() {
	s := "банан"
	m := make(map[rune]int)

	for _, r := range s {
		m[r] += 1
	}

	fmt.Println(m)
}