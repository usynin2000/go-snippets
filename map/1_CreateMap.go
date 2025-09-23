package main

import "fmt"

func main() {
	// Создание map с литералом
	ages := map[string]int{
		"Alice": 25,
		"Bob": 30,
	}

	// Создание пустой  map через make
	m := make(map[string]int)

	//  Добавление элемента
	m["Charlie"] = 40

	fmt.Println("ages: ", ages)
	fmt.Println("m: ", m)

}