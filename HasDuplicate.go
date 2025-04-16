package main

import "fmt"

func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums  {
		if seen[num] {
			return true
		}
		seen[num] = true

	}
	return false
}


func main() {
	nums := []int{1, 2, 3, 4, 2}
	res := hasDuplicate(nums)
	fmt.Printf("Result %t", res)

}