// Create a Rectangle struct with fields W and H (width and height).
// Write a method Area() that returns the area.
// In main create a rectangle 10x5 and print its area.

// Expected output:
package main

import "fmt"

type Rectangle struct { W, H int}

func (r Rectangle) Area() int { return r.H * r.W}

func main() {
	rec := Rectangle{10, 5}

	res := rec.Area()

	fmt.Println(res)
}