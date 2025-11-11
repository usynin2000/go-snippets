package main

import "fmt"

func main() {
	var arr [5]int // array with 5 elements, by default filled with zeros
	arr[0] = 10 
	arr[1] = 20

	fmt.Println("Массив:", arr) // [10, 20, 0, 0 , 0]
	fmt.Println("Длина массива:", len(arr))
}

// 💡 Arrays have a fixed length, cannot be changed.