package main

import "fmt"

func main() {
	// Эмуляция while цикла
	counter := 0
	for counter < 5 {
		fmt.Printf("while цикл: %d\n", counter)
		counter++
	}

	// Бесконечный циклл с условием выхода
	i := 0
	for {
		if i >= 3 {
			break
		} // выход из цикла }
		fmt.Printf("Бесконечный цикл %d\n", i)
		i++
	}

	// Цикл с условием и continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Нечетные числа: %d\n", i)
	}

}
