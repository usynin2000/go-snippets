// 7. Замена

// Создай строку "hello world".
// Замени "world" на "Go".
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"

	s = strings.ReplaceAll(s, "world", "Go")
	fmt.Println(s)
}