package main

import "fmt"

func main() {
	slice := []int{1, 2, 3} // create slice
	fmt.Println("Slice:", slice)


	slice = append(slice, 4, 5) // add elements
	fmt.Println("After append:", slice)
	fmt.Println("Length of slice:", len(slice))
}

// Slice is a dynamic array, can be changed size with append.

// | Characteristic     | Array                       | Slice                            |
// | ------------------ | --------------------------- | -------------------------------- |
// | Size               | Fixed                       | Dynamic                          |
// | Function passing   | Copied fully                | Passed by reference              |
// | Usage              | Rarely used directly        | Practically always used          |
// | Access to elements | Direct, through index        | Direct, through index or slice    |
// | Storage            | Array values                | Structure with pointer to array   |
