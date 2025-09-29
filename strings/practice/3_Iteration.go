// 3. Перебор

// Пройди по строке "Go язык" циклом for range и выведи индекс и символ.

package main

import "fmt"

func main() {
	s := "GO язык"

	for i, r := range s {
		fmt.Println("Index:", i, "Character", string(r))
	}
}