package main

import "fmt"

func main() {
	// Simple var
	x := 10

	fmt.Println("Value x:", x)

	// Getting pointer to the var x
	ptr := &x
	fmt.Println("Address x:", ptr)
	fmt.Println("Value by using pointer", *ptr)

	// Меняем значение через указатель
	// Change the value by using pointer
	*ptr = 20
	fmt.Println("New value x:", x)

	// Передача указателя в функцию
	// Passing pointer to func
	changeValue(ptr)
	fmt.Println("Value of x after func:", x)
}

// Функция, которая изменяет значение переменной по указателю
// Func that change vlue of var by using pointer
func changeValue(p *int) {
	*p = *p + 5
}
