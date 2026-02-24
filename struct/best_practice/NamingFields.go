package main

import "fmt"


type Person struct {
	Name string // when we want to use it in other packages
	age int // only in that package
}

func main() {
	p := Person{"Sergo", 24}
	fmt.Println(p)
}
