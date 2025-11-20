package main

import "fmt"


// Constraint only numbers
type Number interface {
	int | int64 | float64 | float32
}



func Sum[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
	// fmt.Println(Sum[[]string{"a", "b"}]) // []string{…} is not a typecompilerNotAType
}