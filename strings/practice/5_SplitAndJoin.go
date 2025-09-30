// 5. Split и Join

// Раздели строку "go is fun" на слова (через пробел) и выведи.
// А потом склей обратно с разделителем "-".
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "go is fun"

	slice := strings.Split(s, " ")
	fmt.Println(slice)

	s = strings.Join(slice, "-")
	fmt.Println(s)

}