package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25}

	// Доступ по ключу
	fmt.Println("Alice: ", ages["Alice"])

	// Если ключа нет, вернется "нулевое значение" (для int это 0)
	fmt.Println("Bob (нет ключа): ", ages["Bob"])


	// проверка существования ключа:
	val, ok := ages["Bob"]
	if ok {
		fmt.Println("Bob найден:", val)
	} else {
		fmt.Println("Bob отсутствует")
	}
}