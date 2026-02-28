package main

import "fmt"

func main() {
	s := "Go!"
	fmt.Println(s[0]) // 71 (byte, not character)
	fmt.Printf("%c\n", s[0])
}
