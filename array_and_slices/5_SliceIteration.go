package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}

	// by indexes
	for i := 0; i < len(slice); i++ {
		fmt.Println("index: ", i, "value: ", slice[i])
	}

	// by using range
	for i, v := range slice {
		fmt.Println("index: ", i, "value: ", v)
	}


}
// range is a very convenient way to iterate through a slice or an array.
