package main

import "fmt"

func main() {
	// Обычная переменная
	x := 10

	fmt.Println("Значение x:", x)

	// Получаем указатель на переменную
	ptr := &x
	fmt.Println("Адрес x:", ptr)
	fmt.Println("Значение через указатель", *ptr)

	// Меняем значение через указатель
	*ptr = 20
	fmt.Println("Новое значение x:", x)

	// Передача указателя в функцию
	changeValue(ptr)
	fmt.Println("Значение x после функции:", x)
}

// Функция, которая изменяет значение переменной по указателю
func changeValue(p *int) {
	*p = *p + 5
}