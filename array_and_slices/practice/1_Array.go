// 1. Array

// Создай массив из 5 чисел и:

// Заполни его числами от 1 до 5.

// Выведи все элементы через цикл for.

package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, v := range arr {
		fmt.Println(v)
	}
}