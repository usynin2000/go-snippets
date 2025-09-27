package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun"
	words := strings.Split(text, " ") // [go, is, fun]
	fmt.Println(words)

	joined := strings.Join(words, "-")
	fmt.Println(joined) // "go-is-fun"
}