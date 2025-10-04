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

type Circle struct {
	R float64
}

// Метод для круга
func (c Circle) Area() float64 {
	return 3.14 * c.R * c.R
}

func main() {
	shapes := []Shape{
		Rectangle{10, 5},
		Circle{3},
	}

	for _, s := range shapes {
		fmt.Println("Area:", s.Area())
	}
}