package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("apple" == "apple") // true
	fmt.Println("apple" < "banana") // true (comparison by Unicode)
	fmt.Println(strings.EqualFold("Go", "gO")) // true without considering the case
}
