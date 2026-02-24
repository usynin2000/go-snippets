// 🔹 Task 1. Shapes

// Create a Shape interface with the Area() float64 method.
// Create two types:

// Rectangle {W, H float64}

// Circle {R float64}

// Both must implement Shape.
// In main, put them in a []Shape slice and print their areas.

// Expected output (approximately):
// Rectangle area: 20
// Circle area: 28.27

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	H, W float64
}

type Circle struct {
	R float64
}

func (r Rectangle) Area() float64 { return r.H * r.W}
func (c Circle) Area() float64 {return c.R * c.R * 3.14}

func main() {
	shapes := []Shape{
		Rectangle{4, 5},
		Circle{3},
	}

	fmt.Println("Rectangle area:", shapes[0].Area())
	fmt.Println("Circle area:", shapes[1].Area())

}