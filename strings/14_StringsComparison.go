package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("apple" == "apple") // true
	fmt.Println("apple" < "banana") // true (сравнение по Unicode)
	fmt.Println(strings.EqualFold("Go", "gO")) // true без учета регистра
}