package main

import "fmt"

func main() {
	n := 1
	for ok := true; ok; ok = n < 5 {
		fmt.Println("n =", n)
		n++
	}
}