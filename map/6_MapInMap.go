package main

import "fmt"


func main() {
	scores := map[string]map[string]int{
		"group1": {"Alice": 90, "Bob": 80},
	}

	// Обновляем оценку
	scores["group1"]["Alice"] = 95

	fmt.Println(scores)
}