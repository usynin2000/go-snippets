package main

import "fmt"

type Rectangle struct {
	Width int
	Height int
}

// метод для Rectangle
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{5, 3}
	fmt.Println("Площадь:", r.Area())
}