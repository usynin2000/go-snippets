package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{3, 4}

	fmt.Println(p1 == p2) // true
	fmt.Println(p1 == p3) // false
}

// Можно сравнивать, если все поля поддерживают сравнение.