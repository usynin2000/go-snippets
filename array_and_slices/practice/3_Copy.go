// Есть срез:

// src := []int{10, 20, 30, 40}


// Сделай независимую копию (так, чтобы при изменении src копия не изменилась) и докажи это выводом на экран.

package main

import "fmt"

func main() {
	src := []int{10, 20, 30, 40}

	cp := make([]int, len(src))

	copy(cp, src)

	fmt.Println(src, cp)
}