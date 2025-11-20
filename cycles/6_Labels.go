package main

import "fmt"

func main() {
	// Использование меток для выхода из вложенных циклов
	fmt.Println("Поиск с использованием меток:")

outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1{
				fmt.Printf("Найдена позиция [%d][%d], выходим из всех циклов\n", i, j)
				break outerLoop
			}
			fmt.Printf("Проверяем [%d][%d] \n", i, j)
		}
	}

	fmt.Println("\nПродолжение с меткой:")

continueLoop:
	for i := 0; i< 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 1 {
				fmt.Printf("Пропускаем итерацию [%d][%d]\n", i, j)
				continue continueLoop
			}
			fmt.Printf("Обрабатываем [%d][%d]\n", i, j)
		}
	}


}

