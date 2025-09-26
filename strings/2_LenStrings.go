package main

import "fmt"

func main() {
	s := "Привет"
	fmt.Println(len(s)) // 12 (каждая кириллическая буква занимает 2 байта в utf-8)
}