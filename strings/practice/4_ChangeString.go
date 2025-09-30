// 4. Изменение строки

// Создай строку "молоко".
// Замени первую букву на 'х' и выведи результат.
package main

import "fmt"

func main() {
	s := "молоко"
	runes := []rune(s)

	runes[0] = 'x'
	s = string(runes)
	fmt.Println(s)

}