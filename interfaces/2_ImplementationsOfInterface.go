package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

// Structure
type Rectangle struct {
	W, H float64
}

// Method for Rectangle
func (r Rectangle) Area() float64{
	return r.H * r.W
}

type Circle struct {
	R float64
}

// Method for Circle
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