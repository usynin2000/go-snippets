package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	part1 := arr[1:3] // elements from index 1 to 2 -> [20, 30]
	part2 := arr[:2] // from start to index 2 (not including) -> [10, 20]
	part3 := arr[2:] // from index 2 to the end [30, 40, 50]

	fmt.Println(part1, part2, part3)
}
