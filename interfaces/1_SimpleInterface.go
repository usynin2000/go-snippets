// 1. Что такое интерфейс?

// Интерфейс в Go — это набор методов.
// Тип «реализует» интерфейс, если у него есть все эти методы (никаких ключевых слов implements не нужно).
package main

import "fmt"

// Интерфейс
type Shape interface {
	Area() float64
}

// Структура
type Rectangle struct {
	W, H float64
}

// Метод для Rectangle
func (r Rectangle) Area() float64{
	return r.H * r.W
}

func main() {
	var s Shape  // переменная типа интерфейса
	s = Rectangle{10, 5} // Rectangle реализует Shape автоматически
	fmt.Println(s.Area())
}