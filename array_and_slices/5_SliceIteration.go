package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}

	// по индексам
	for i := 0; i < len(slice); i++ {
		fmt.Println("index: ", i, "value: ", slice[i])
	}

	// через range
	for i, v := range slice {
		fmt.Println("index: ", i, "value: ", v)
	}


}

//💡 range — очень удобный способ пройтись по срезу или массиву.