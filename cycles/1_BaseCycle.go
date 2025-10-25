package main

import "fmt"

func main() {
	// Простой цикл for
	for i := 0; i < 5; i++ {
		fmt.Printf("Итерация %d\n", i)
	}

	// Цикл с декрементами
	for i := 10; i > 0; i-- {
		fmt.Printf("Обртный отсчет: %d\n", i)
	}

	// Цикл с шагом
	for i := 0; i <20; i += 2 {
		fmt.Printf("Четные числа: %d\n", i )
	}

}