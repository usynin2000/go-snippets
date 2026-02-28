package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s) // string → bytes
	fmt.Println(b) // [104 101 108 108 111]

	s2 := string(b) // bytes → string
	fmt.Println(s2) // hello
}
