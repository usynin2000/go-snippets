package main


import "fmt"

func main() {
	s := "Go язык"
	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}
}
