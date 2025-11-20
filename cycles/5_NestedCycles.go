package main

import "fmt"


func main() {
	// Таблица умножения
	fmt.Println("Таблица умножения")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}

	// Поиск в двумерном слайсе
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	target := 5
	found := false

	for i, row := range matrix {
		for j, value := range row {
			if value == target {
				fmt.Printf("Найдено %d в позици [%d][%d]\n", target, i, j)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

}