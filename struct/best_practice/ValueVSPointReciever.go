package main

import "fmt"

type Counter struct { Value int}


//  Изменяем структуру -> pointer receiver
func (c *Counter) Inc() {c.Value++}

// Только читаем -> value receiver
func (c Counter) Get() int { return c.Value}


func main() {

	c := Counter{0}

	pointerC := &c

	pointerC.Inc()
	pointerC.Inc()
	pointerC.Inc()

	fmt.Println(c.Get())

}

// Value receiver (func (r Rectangle))
	// Используется, если метод не изменяет структуру.
	// Работает для маленьких структур (не копируем много данных).


// Pointer receiver (func (r *Rectangle))
	// Используется, если метод изменяет структуру.
	// Экономит память для больших структур.