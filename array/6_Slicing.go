package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	part1 := arr[1:3] // элементы с индекса 1 по 2 -> [20, 30]
	part2 := arr[:2] // с начала до индекса 2 (не включая) -> [10, 20]
	part3 := arr[2:] //с индекса 2 и до конца [30, 40, 50]

	fmt.Println(part1, part2, part3)
}