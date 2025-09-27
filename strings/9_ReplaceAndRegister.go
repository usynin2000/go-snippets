package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Println(strings.ReplaceAll(s, "world", "go")) // hello go
	fmt.Println(strings.ToUpper(s)) // HELLO WORLD
	fmt.Println(strings.ToLower("ПрИвЕт")) // привет
}