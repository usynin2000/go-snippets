package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}

	// Обновление значения
	ages["Alice"] = 26

	// Удалние 
	delete(ages, "Bob")

	fmt.Println(ages)
}