package main

import "fmt"


func main() {
	var ptr * int // указатель но без адреса -> nil

	if ptr == nil {
		fmt.Println("Указатель пустой")
	}

	x := 7

	ptr = &x

	fmt.Println("Теперь ptr указывает на:", *ptr)
}

// 👉 Указатель по умолчанию равен nil, пока ему не присвоен адрес.