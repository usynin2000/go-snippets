package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s) // строка → байты
	fmt.Println(b) // [104 101 108 108 111]

	s2 := string(b) // байты → строка
	fmt.Println(s2) // hello
}