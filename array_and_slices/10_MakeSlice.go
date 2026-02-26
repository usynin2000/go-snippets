package main

import "fmt"

func main() {
	slice := make([]int, 3, 5) // length 3, capacity 5)
	fmt.Println(slice) // [0 0 0]
	fmt.Println("len: ", len(slice), "cap:", cap(slice))
}

// 💡 In Go slices have length (len) and capacity (cap).
