package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 40}
	fmt.Println("Количество элементов: ", len(ages))
}