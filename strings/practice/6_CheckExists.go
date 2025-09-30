// 6. Проверки

// Создай строку "golang is cool".
// Проверь:

// содержит ли она "go",

// начинается ли с "golang",

// заканчивается ли на "cool".

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "golang is cool"


	fmt.Println(strings.Contains(s, "go"))
	fmt.Println(strings.HasPrefix(s, "golang"))
	fmt.Println(strings.HasSuffix(s, "cool"))

}