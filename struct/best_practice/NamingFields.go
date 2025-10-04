package main

import "fmt"


type Person struct {
	Name string // доступно из других пакектов
	age int // только внутри пакета
}

func main() {
	p := Person{"Sergo", 24}
	fmt.Println(p)
}