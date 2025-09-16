package main

import "fmt"


// Изменение перменной по указателю
func increment(n *int) {
	*n++
}

// Фукнция, возвращая указатель
func createPointer(val int) *int {
	return &val
}

func main() {
	// 1. Базовый пример
	x := 10
	ptr := &x  // берем адрес переменной
	fmt.Println("1) x:", x, "| адрес:", ptr, "| через указатель:", *ptr)

	*ptr = 20 // меняем через указатель
	fmt.Println("новое значение x:", x)

	// 2. Указатель как аргумент функции
	increment(&x)
	fmt.Println("2) после increment:", x)

	// 3. Указатель на указатель
	p1 := &x
	p2 := &p1

	**p2 = 42
	fmt.Println("3) через двойной указатель", x)

	// 4. Указатель, возвращаемый из функции
	p3 := createPointer(99)
	fmt.Println("4) createPointer -> ", *p3)

	// 5. Nil-указатель
	var p4 *int // пока пустой
	if p4 == nil {
		fmt.Println("5) p4 пустой (nil)")
	}
	y := 7
	p4 = &y
	fmt.Println("  теперь p4 указывает на:", *p4)

	// 6. Указатель на структуру
	type Person struct {
		Name string
		Age int
	}
	p := Person{Name: "Alex", Age: 24}
	pPtr := &p
	pPtr.Age = 25 // можно менять поля без (*pPtr)
	fmt.Println("6) стрктура через указыватель:", p)
}


// 📌 В этом сниппете:
// Базовое использование указателя (&, *).
// Изменение переменной через функцию.
// Указатель на указатель.
// Возврат указателя из функции.
// Nil-указатель.
// Указатель на структуру.