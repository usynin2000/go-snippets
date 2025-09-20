package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)) // создаем новый срез такой же длины

	copy(b, a)  // копируем данные
	b[0] = 100

	fmt.Println("a: ", a) // [1, 2, 3]
	fmt.Println("b: ", b) // [100, 2, 3]
}

// 💡 copy создаёт новый срез с отдельной памятью.