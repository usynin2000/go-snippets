package main

import "fmt"

func main() {
	s := "Go!"
	fmt.Println(s[0]) // 71 (байт, а не символ)
	fmt.Printf("%c\n", s[0])
}