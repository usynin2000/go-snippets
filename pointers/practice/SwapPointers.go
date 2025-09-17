package main

import "fmt"


func swap (a *int, b *int){
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 3
	b := 7

	swap(&a, &b)

	fmt.Println(a, b)
}