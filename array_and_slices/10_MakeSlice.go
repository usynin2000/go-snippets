package main

import "fmt"

func main() {
	slice := make([]int, 3, 5) // длина 3, вместимость 5)
	fmt.Println(slice) // [0 0 0]
	fmt.Println("len: ", len(slice), "cap:", cap(slice))
}

// 💡 В Go у срезов есть длина (len) и вместимость (cap).