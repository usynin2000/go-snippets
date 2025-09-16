package main

import "fmt"

func createPointer(val int) *int {
	return &val
}

func main() {
	ptr := createPointer(99)
	fmt.Println("Значение", *ptr)
}

// 👉 Функция может возвращать указатель.
//  Go сам гарантирует, что память не "исчезнет", даже если переменная была локальной в функции.