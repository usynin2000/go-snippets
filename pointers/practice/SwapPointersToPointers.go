package main

import "fmt"

func swapPointers(a **int, b **int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 5
	b := 10

	pa := &a
	pb := &b

	swapPointers(&pa, &pb)

	fmt.Println(*pa, *pb)
}