// 🔹 Задание 2. Вложенные структуры

// Создай структуру Person с полями Name и Address.
// Структура Address должна содержать поля City и Zip.
// Создай человека "Alice" из города "Moscow" с индексом "101000" и выведи его имя и город.

// Ожидаемый вывод:

// Alice lives in Moscow

package main

import "fmt"


type Address struct {
	City string
	Zip int
}

type Person struct {
	Name string
	Address
}

func main() {
	p := Person{"Alice", Address{"Moscow", 123432}}

	fmt.Println(p.Name, "lives in", p.City)
}