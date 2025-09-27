package main

import "fmt"

func main() {
	text := `Это многострочная строка.
Она сохраняет переносы
 и пробелы.
	`
	fmt.Println(text)
}