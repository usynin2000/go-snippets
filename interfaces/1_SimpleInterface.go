// 1. What is an interface?

// Interface in Go is a set of methods.
// Type «implements» interface if it has all these methods (no keywords implements needed).

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

func main() {
	var s Shape  // var of interface type
	s = Rectangle{10, 5} // Rectangle implements Shape automatically
	fmt.Println(s.Area())
}
