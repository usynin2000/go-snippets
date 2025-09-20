package main

import "fmt"


func main() {
	slice := []int{1, 2, 3, 4 ,5}

	index := 2 // удалим элемент с индексом 2 (число 3)

	slice = append(slice[:index], slice[index + 1:]...) // ... берёт слайс и "раскладывает" его элементы как отдельные аргументы функции.

	fmt.Println(slice) // [1, 2, 4, 5]
}

// slice[:index] — все элементы до удаляемого.

// slice[index+1:] — все элементы после удаляемого.

// append(..., ......) — склеиваем их в новый слайс.


// В Go ... в таком контексте называется spread-оператор (или "variadic expansion").