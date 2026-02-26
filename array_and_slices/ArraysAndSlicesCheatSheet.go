package main

import "fmt"

func main() {
	// 1. Creating slices
	s1 := []int{1, 2, 3} // literal
	s2 := make([]int, 3, 5)
	fmt.Println("1) s1: ", s1, "| s2: ", s2, "len:", len(s2), "cap:", cap(s2))

	// 2.Slicing
	arr := []int{10, 20, 30, 40, 50}
	part1 := arr[1:3] // indexes 1 and 2
	part2 := arr[:2] // from start to index 2
	part3 := arr[2:]  // from 2 to the end
	fmt.Println("2) Slicing: ", part1, part2, part3)

	// 3. Adding an element
	s1 = append(s1, 4, 5)
	fmt.Println("3) append:", s1)

	// 4. Concatenation of slices
	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...)
	fmt.Println("4) Concatenation:", a)

	// 5. Copying a slice
	src := []int{1, 2, 3}
	cp_src := make([]int, len(src))
	copy(cp_src, src)
	cp_src[0] = 100
	fmt.Println("5) copy: src =", src, "| cp_src =", cp_src)

	// 6. Removing an element
	nums := []int{1, 2, 3, 4, 5}
	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("6) Removing an element: ", nums)

	// 7. Iteration
	for i, v := range s1 {
		fmt.Println("7) range:", i, v)
	}

	// 8. nil and empty slices
	var sNil []int // equals nil, len == 0
	sEmpty := []int{} // not equals nil, although len == 0
	fmt.Println("8) nil slice: ", sNil, "len:", len(sNil), "is_nil?", sNil == nil)
	fmt.Println("8) empty slice:", sEmpty, "len:", len(sNil), "is_nil?", sEmpty == nil)

	// 9. Filters
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	var short []string
	for _, name := range names {
		if len(name) <= 4 {
			short = append(short, name)
		}
	}
	fmt.Println("9) Filtering:", short)

	// 10. Memory allocation: aliasing
	orig := []int{1, 2, 3, 4, 5}
	alias := orig[:3] // points to the same array
	alias[0] = 99
	fmt.Println("10) aliasing:", "orig =", orig, "| alias =", alias)

	// To avoid: copy
	independent := append([]int(nil), orig[:3]...)
	independent[0] = 123
	fmt.Println("   independent copy:", "orig =", orig, "| independent =", independent)

}
