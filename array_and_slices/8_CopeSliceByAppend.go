package main

import "fmt"


func main() {
	a := []int{1, 2, 3}
	b := append([]int(nil), a...) //создаем копию,  часто считается более «go-шным» при клонировании.

	b[0] = 100

	fmt.Println("a: ", a) // [1 2 3]
	fmt.Println("b: ", b) // [100, 2, 3]
}