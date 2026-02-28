package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "golang is cool"

	fmt.Println(strings.Contains(s, "go")) // true
	fmt.Println(strings.HasPrefix(s, "go")) // true
	fmt.Println(strings.HasSuffix(s, "cool")) // true
	fmt.Println(strings.Index(s, "lang")) // 2 (position in bytes)
}
