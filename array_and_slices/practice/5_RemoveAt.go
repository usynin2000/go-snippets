// Удаление с сохранением порядка

// Напиши функцию removeAt, которая удаляет элемент по индексу и сохраняет порядок элементов:
// nums := []int{1, 2, 3, 4, 5}
// nums = removeAt(nums, 2)
// fmt.Println(nums) // [1 2 4 5]


package main

import "fmt"

func removeAt(lst []int, index int) []int {
	result := append(lst[:index], lst[index +1:]...)
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = removeAt(nums, 2)

	fmt.Println(nums) // [1 2 4 5]
}