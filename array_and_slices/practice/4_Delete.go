// 4. Удаление

// Из среза:

// nums := []int{1, 2, 3, 4, 5}


// удали число 3.

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	index := 2

	nums = append(nums[:index], nums[index + 1:]...)

	fmt.Println(nums)
}