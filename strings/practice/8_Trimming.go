// 8. Trim

// Создай строку " hello world ".
// Удалите пробелы по краям.
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world  "
	s = strings.Trim(s, " ")
	fmt.Println(s)
}