package main

import "fmt"


func main() {
	// Можно создать структуру без объявления типа

	p := struct {
		Name string
		Age int
	}{
		Name: "Charlie",
		Age: 22,
	}

	fmt.Println(p)
}