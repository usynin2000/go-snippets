package main

import "fmt"

func increase(a *int) {
	*a = *a + 1
}

func main() {
	x := 10

	increase(&x)

	fmt.Println(x)
}