package main

import "fmt"

func main() {
	s := "Привет"
	fmt.Println(len(s)) // 12 (each russian letter takes 2 bytes in utf-8)
}
