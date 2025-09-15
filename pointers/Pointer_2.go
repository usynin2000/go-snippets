package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// Создаем переменную типа Person
	p := Person{Name: "Serega", Age: 24}

	// Берем указатель на структуру
	ptr := &p

	// Меняем поле через указатель
	ptr.Age = 150

	fmt.Println("Имя:", p.Name)
	fmt.Println("Возраст", p.Age)

}

// ВЫВОД: 
// 👉 Поля структуры можно менять через указатель без явного разыменования 
// ((*ptr).Age эквивалентно ptr.Age).
