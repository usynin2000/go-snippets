package main

import "fmt"

type Person struct {
	Name string
	Age int
}


func main() {
	// Полная инициализация
	p1 := Person{Name: "Alice", Age: 25}

	// Без указания полей (подрядок, обязателен)
	p2 := Person{"Bob", 30}

	// Пустая структура (значения по умолчанию)
	p3 := Person{}

	fmt.Println(p1) // {Alice 25}
	fmt.Println(p2) // {Bob 30}
	fmt.Println(p3) //{ 0}
}