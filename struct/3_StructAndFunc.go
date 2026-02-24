package main

import "fmt"

type Rectangle struct {
	Width int
	Height int
}

// Func that recieve as papam struct
func Area (r Rectangle) int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 5, Height: 3}
	fmt.Println("Area:", Area(r)) // 15
}
