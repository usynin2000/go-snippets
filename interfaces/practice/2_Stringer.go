// 🔹 Задание 2. Stringer

// Создай структуру Person {Name string; Age int}.
// Сделай так, чтобы при печати через fmt.Println(p) выводилось:

// Alice (25 years old)
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
	p := Person{"Alice", 25}

	fmt.Println(p)
}