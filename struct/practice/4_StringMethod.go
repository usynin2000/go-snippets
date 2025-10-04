// 🔹 Метод String()

// Создай структуру Person с полями Name и Age.
// Реализуй для неё метод String() string, который возвращает строку вида:
// Name: Alice, Age: 25
// И выведи через fmt.Println(p).
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	p := Person{"Alice", 25}
	fmt.Println(p.String())
	// Это работает, но в Go есть фишка: если структура реализует метод String() string, 
	// то fmt.Println(p) вызывает его автоматически 🎉
	fmt.Println(p)

}