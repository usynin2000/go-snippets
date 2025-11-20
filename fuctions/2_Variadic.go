package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}



func join(sep string, parts ...string) string { // Variadic can be only last param
	res := ""
	for i, p := range parts {
		if i > 0{
			res += sep
		}
		res += p
	}
	return res
}

func main() {
	fmt.Println("sum: ", sum(1, 2, 3))
	s := []int{4, 5, 6}
	fmt.Println("sum(slice..):", sum(s...)) // ... disclose slice!!
	fmt.Println("join:", join("-", "a", "b", "c"))
}