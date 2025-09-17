package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func haveBirthday(obj *Person) {
	obj.Age = obj.Age + 1
}

func main() {
	p := Person{Name: "Alex", Age: 25}

	haveBirthday(&p)
	fmt.Println(p.Age)
}