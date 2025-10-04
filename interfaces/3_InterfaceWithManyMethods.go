package main

import "fmt"

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct{}
func (Dog) Speak() string {return "Woof"}
func (Dog) Move() string {return "Runs"}

type Cat struct{}
func (Cat) Speak() string {return "Meow"}
func (Cat) Move() string {return "Sneaks"}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println("Dog", a.Speak(), a.Move())

	a = Cat{}
	fmt.Println("Cat:", a.Speak(), a.Move())
}
