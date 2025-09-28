package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "     hello world   "
	fmt.Println(strings.TrimSpace(s)) // hello world
	fmt.Println(strings.Trim(s, " hd")) // ello worl
}