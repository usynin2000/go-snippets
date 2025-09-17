package main

import "fmt"

func main() {
	slice := []int{1, 2, 3} // создаем срез
	fmt.Println("Срез", slice)


	slice = append(slice, 4, 5) // добавляем элмементы
	fmt.Println("После append:", slice)
	fmt.Println("Длина срезы:", len(slice))
}

// 💡 Срезы — это динамические массивы, можно изменять размер с помощью append.