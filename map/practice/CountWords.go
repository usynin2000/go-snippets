// 4. Подсчёт слов

// Дана строка:

// text := "go is fun and go is fast"


// Нужно посчитать, сколько раз встречается каждое слово, и сохранить в map[string]int.
// Выведи результат.
package main

import (
	"fmt"
	"strings"
)

func main() {

	counterWords := make(map[string]int)

	text := "go is fun and go is fast"
	words := strings.Split(text, " ")


	for _, word := range words {
		_, ok := counterWords[word]
		if ok{
			counterWords[word]++
		} else {
			counterWords[word] = 1
		}
		fmt.Println(word)
	}

	fmt.Println(counterWords)

}
