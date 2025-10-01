package main

import "fmt"

type Rectangle struct {
	Width int
	Height int
}

// Функция, которая принимает структуру
func Area (r Rectangle) int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 5, Height: 3}
	fmt.Println("Площадь:", Area(r)) // 15
}